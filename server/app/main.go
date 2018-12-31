package main

import (
	"context"
	"errors"
	"html/template"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/project-hermes/hermes-app/server/graph"
	"google.golang.org/api/option"

	"github.com/project-hermes/hermes-app/server/model"
	"github.com/project-hermes/hermes-app/server/wrapper"
)

const (
	// FirebaseApp is a constant for pulling app from gin
	FirebaseApp = "FirebaseApp"

	// FirebaseAuth is a constant for pulling auth from gin
	FirebaseAuth = "FirebaseAuth"
)

type templateParams struct {
	Notice string
	Name   string
}

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

func init() {

}

func main() {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	router.Use(InjectFirebase)

	appGroup := router.Group("/")
	appGroup.GET("/", helloForm)

	gqlPlayground := gin.WrapH(handler.Playground("GraphQL playground", "/query"))
	router.GET("/query", gqlPlayground)

	dbClient, _ := wrapper.NewClient(context.Background(), "project-hermes-staging")
	diveInt := model.NewDiveImplementation(dbClient)
	sensorInt := model.NewSensorImplementation(dbClient)

	resolver := graph.NewResolver(diveInt, sensorInt)
	gql := gin.WrapH(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver})))
	router.POST("/query", gql)

	router.Run()
}

func helloForm(c *gin.Context) {
	params := templateParams{}
	indexTemplate.Execute(c.Writer, params)
	return
}

// InjectFirebase will inject firebase
func InjectFirebase(c *gin.Context) {
	log.Printf("Loading firebase app")
	c.Keys = map[string]interface{}{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	opt := option.WithHTTPClient(http.DefaultClient)
	opt = option.WithScopes("https://www.googleapis.com/auth/cloud-platform", "https://www.googleapis.com/auth/userinfo.email")
	databaseURL := "https://project-hermes-staging.firebaseio.com"
	config := &firebase.Config{
		DatabaseURL: databaseURL,
		ProjectID:   "project-hermes-staging",
	}

	if app, err := firebase.NewApp(ctx, config, opt); err != nil {
		log.Fatalf("Unable to load firebase app: %s", err)
		c.AbortWithError(http.StatusServiceUnavailable, errors.New("an error occurred while processing your request"))
	} else {
		c.Keys[FirebaseApp] = app
		c.Next()
	}
}

// InjectFirebaseAuthClient will inject the firebase auth
func InjectFirebaseAuthClient(c *gin.Context) {
	log.Printf("Loading firebase auth client")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if app, ok := c.Keys[FirebaseApp].(*firebase.App); !ok {
		log.Fatalf("Could not get firebase app")
	} else if authClient, err := app.Auth(ctx); err != nil {
		log.Fatalf("Unable to create auth client: %v", err)
	} else {
		c.Keys[FirebaseAuth] = authClient
		c.Next()
	}
}

func verifyIDToken(c *gin.Context, app *firebase.App, idToken string) *auth.Token {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if authClient, ok := c.Keys[FirebaseAuth].(*auth.Client); !ok {
		log.Fatalf("Could not get firebase app")
		c.AbortWithError(http.StatusInternalServerError, errors.New("an issue occurred on the server"))
		return nil
	} else if token, err := authClient.VerifyIDToken(ctx, idToken); err != nil {
		log.Fatalf("error verifying ID token: %v", err)
		c.AbortWithError(http.StatusForbidden, errors.New("unauthorized"))
		return nil
	} else {
		log.Printf("Verified ID Token: %v", token)
		return token
	}
}

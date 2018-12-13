package main

import (
	"context"
	"errors"
	"html/template"
	"net/http"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"github.com/99designs/gqlgen/handler"
	"github.com/project-hermes/hermes-app/server/graph"

	"github.com/project-hermes/hermes-app/server/model"

)

const (
	// FirebaseApp is a constant for pulling app from gin
	FirebaseApp = "FirebaseApp"
)

type templateParams struct {
	Notice string
	Name string
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

	fsClient, _ := firestore.NewClient(context.Background(), "project-hermes-staging")
	diveInt := model.NewDiveImplementation(fsClient)

	resolver := graph.NewResolver(diveInt)
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
		ProjectID: "project-hermes-staging",
	}

	if app, err := firebase.NewApp(ctx, config, opt); err != nil {
		log.Fatalf("Unable to load firebase app: %s", err)
		c.AbortWithError(http.StatusServiceUnavailable, errors.New("an error occurred while processing your request"))
	}  else {
		c.Keys[FirebaseApp] = app
		c.Next()
	}
}
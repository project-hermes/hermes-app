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
	"github.com/gin-gonic/gin/binding"
	"github.com/project-hermes/hermes-app/server/graph"
	"github.com/project-hermes/hermes-app/server/model"
	"github.com/project-hermes/hermes-app/server/protobuf"
	"github.com/project-hermes/hermes-app/server/wrapper"
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
	router.Use(injectFirebase)
	router.Use(injectModel)

	appGroup := router.Group("/")
	appGroup.GET("/", helloForm)

	router.POST("/dive", createDive)



	// dbClient, _ := wrapper.NewClient(context.Background(), "project-hermes-staging")
	// diveInt := model.NewDiveImplementation(dbClient)
	// sensorInt := model.NewSensorImplementation(dbClient)

	// resolver := graph.NewResolver(diveInt, sensorInt)
	// gql := gin.WrapH(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver})))
	// router.POST("/query", gql)

	if err := router.Run(); err != nil {
		log.Fatalf("unable to start gin router %s", err.Error())
	}
}

func helloForm(c *gin.Context) {
	params := templateParams{}
	indexTemplate.Execute(c.Writer, params)
	return
}

// TODO: Clean this up and prepare for it to be called by PubSub
func createDive(c *gin.Context) {
	dive := &protobuf.Dive{}
	err := binding.ProtoBuf.Bind(c.Request, dive)

	if err != nil {
		log.Fatalf("there was an issue with the dive protobuf binding %s", err)
		c.AbortWithError(http.StatusBadRequest, errors.New("unable to parse protobuf dive"))
	}

	if diveClient, ok := c.Keys[Dive].(model.DiveInterface); !ok {
		log.Fatalf("Could not get dive implementation from create dive")
	} else if err := diveClient.Create(c, model.NewDive(dive)); err != nil {

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

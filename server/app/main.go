package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	"github.com/project-hermes/hermes-app/server/graph"
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

	appGroup := router.Group("/")
	appGroup.GET("/", helloForm)

	gqlPlayground := gin.WrapH(handler.Playground("GraphQL playground", "/query"))
	router.GET("/query", gqlPlayground)

	gql := gin.WrapH(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))
	router.POST("/query", gql)
	
	router.Run()
}

func helloForm(c *gin.Context) {
	params := templateParams{}
	indexTemplate.Execute(c.Writer, params)
	return
}
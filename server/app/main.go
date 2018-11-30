package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
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
	
	router.Run()
}

func helloForm(c *gin.Context) {
	params := templateParams{}
	indexTemplate.Execute(c.Writer, params)
	return
}
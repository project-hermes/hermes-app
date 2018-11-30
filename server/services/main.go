package main

import (
	"fmt"
	"html/template"
	"net/http"

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
	appGroup.POST("/", basicPostForm)
	
	router.Run()
}

func helloForm(c *gin.Context) {
	params := templateParams{}
	indexTemplate.Execute(c.Writer, params)
	return
}

func basicPostForm(c *gin.Context) {
	params := templateParams{}
	name := c.Request.FormValue("name")
	if name == "" {
		name = "Anonymous Gopher"
	} else {
		params.Name = name
	}

	if c.Request.FormValue("message") == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		params.Notice = "No message provided"
		indexTemplate.Execute(c.Writer, params)
		return
	}

	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", name)
	indexTemplate.Execute(c.Writer, params)
}
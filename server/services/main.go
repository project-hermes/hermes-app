package services

import (
	"google.golang.org/appengine"
	"hermes/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	server.ApiEntry(router)
	http.Handle("/", router)

	appengine.Main()
}

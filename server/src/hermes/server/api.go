package server

import (
	"github.com/gin-gonic/gin"
	"hermes/middleware"
)

func ApiEntry(router *gin.Engine) {
	appApi := router.Group("/api")
	appApi.Use(SetupSession(), ApiInjector())
	{

		diveGroup := appApi.Group("/dive")
		diveGroup.GET("/", middleware.TestDive)
		diveGroup.GET("/:id", middleware.DiveGet)
		diveGroup.POST("/", middleware.DiveCreate)

		userGroup := appApi.Group("/user")
		userGroup.GET("/:uid", middleware.GetUser)
	}
}

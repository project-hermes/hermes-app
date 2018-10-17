package server

import (
	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"hermes/model"
	"net/http"
)

func ApiInjector() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		var firebaseApp *firebase.App
		if appengine.IsDevAppServer() {
			tempApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("/Users/bradfordbazemore/Devel/hermes/hermes-app-engine/project-hermes-staging-firebase-adminsdk-q2yxf-fd6ecd39e6.json"))
			if err != nil {
				log.Errorf(ctx, err.Error())
				c.AbortWithStatus(http.StatusNotFound)
			} else {
				firebaseApp = tempApp
			}
		} else {
			tempApp, err := firebase.NewApp(ctx, nil)
			if err != nil {
				log.Errorf(ctx, err.Error())
				c.AbortWithStatus(http.StatusForbidden)
			} else {
				firebaseApp = tempApp
			}
		}
		c.Set("UserInterface", model.NewUserImplementation(ctx, firebaseApp))
		c.Set("DiveInterface", model.NewDiveImplementation(ctx, firebaseApp))
	}
}

func SetupSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		ctx := appengine.NewContext(c.Request)
		var app *firebase.App
		if appengine.IsDevAppServer() {
			tempApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("/Users/bradfordbazemore/Devel/hermes/hermes-app-engine/project-hermes-staging-firebase-adminsdk-q2yxf-fd6ecd39e6.json"))
			if err != nil {
				log.Errorf(ctx, err.Error())
				c.AbortWithStatus(http.StatusNotFound)
			} else {
				app = tempApp
			}
		} else {
			tempApp, err := firebase.NewApp(ctx, nil)
			if err != nil {
				log.Errorf(ctx, err.Error())
				c.AbortWithStatus(http.StatusForbidden)
			} else {
				app = tempApp
			}
		}
		if app == nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
		client, err := app.Auth(ctx)
		if err != nil {
			log.Errorf(ctx, err.Error())
			c.AbortWithStatus(http.StatusForbidden)
		}
		_, err = client.VerifyIDTokenAndCheckRevoked(ctx, token)
		if err != nil {
			log.Errorf(ctx, err.Error())
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

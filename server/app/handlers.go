package main

import (
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/project-hermes/hermes-app/server/model"
	"github.com/project-hermes/hermes-app/server/wrapper"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

const (
	// FirebaseApp is a constant for pulling app from gin
	FirebaseApp = "FirebaseApp"

	// FirebaseAuth is a constant for pulling auth from gin
	FirebaseAuth = "FirebaseAuth"

	// Dive is a constant for pulling the dive client from gin
	Dive = "DiveImplementation"
)

func injectModel(c *gin.Context) {
	dbClient, _ := wrapper.NewClient(context.Background(), "project-hermes-staging")
	diveInt := model.NewDiveImplementation(dbClient)
	c.Keys[Dive] = diveInt
}

// injectFirebase will inject firebase
func injectFirebase(c *gin.Context) {
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
package middleware

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func GetUser(c *gin.Context) {
	if uid := c.Param("uid"); uid == "" {
		c.Status(http.StatusNotFound)
		return
	} else if userInt, found := c.Keys["UserInterface"].(model.UserInterface); !found {
		c.Status(http.StatusNotFound)
		return
	} else if user, err := userInt.Get(uid); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUpdateUser(c *gin.Context) {
	if uid := c.Param("uid"); uid == "" {
		c.Status(http.StatusNotFound)
		return
	} else if userInt, found := c.Keys["UserInterface"].(model.UserInterface); !found {
		c.Status(http.StatusNotFound)
		return
	} else if user, err := userInt.CreateUpdateFromFirebase(uid); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, user)
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/datastore"

	"github.com/project-hermes/hermes-app/server/hermes/model"
	"net/http"
	"time"
)

func DiveGet(c *gin.Context) {
	if key, err := datastore.DecodeKey(c.Param("id")); err != nil {
		c.String(http.StatusBadRequest, "")
		return
	} else {
		if diveInt, found := c.Keys["DiveInterface"].(model.DiveInterface); !found {
			c.String(http.StatusNotFound, "")
			return
		} else {
			if dive, err := diveInt.Get(key); err != nil {
				c.String(http.StatusNotFound, "")
				return
			} else {
				c.JSON(http.StatusOK, diveInt.Public(dive))
			}
		}
	}
}

func DiveCreate(c *gin.Context) {
	var tempPublicDive model.DivePublic
	if err := c.BindJSON(&tempPublicDive); err != nil {
		c.String(http.StatusBadRequest, "%s", err)
		return
	} else {
		if diveInt, found := c.Keys["DiveInterface"].(model.DiveInterface); !found {
			c.String(http.StatusNotFound, "")
			return
		} else {
			if key, err := diveInt.Create(diveInt.Private(tempPublicDive), time.Now()); err != nil {
				c.String(http.StatusBadRequest, "%s", err)
				return
			} else {
				c.String(http.StatusOK, key.Encode())
				return
			}
		}
	}
}

func TestDive(c *gin.Context) {
	if diveInt, found := c.Keys["DiveInterface"].(model.DiveInterface); !found {
		c.String(http.StatusNotFound, "")
		return
	} else {
		if data, err := diveInt.GetAggregation(); err != nil {
			c.Status(http.StatusNotFound)
			return
		} else {
			c.Data(http.StatusOK, "text/plain", data)
		}
	}
}

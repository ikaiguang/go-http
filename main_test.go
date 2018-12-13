package gohttp

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNew(t *testing.T) {
	engine := New()

	var handler = func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}

	RegisterRoutes(engine, []*Route{NewRoute("GET", "ping", handler)})

	if err := Run(engine, ":8081"); err != nil {
		panic(err)
	}
}

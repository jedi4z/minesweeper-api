package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestAdapter struct {
}

func NewRestEngine() *gin.Engine {
	r := gin.Default()
	s := &RestAdapter{}

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", s.pingHandler)
	}

	return r
}

func (r RestAdapter) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

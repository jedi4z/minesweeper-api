package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/minesweeper-api/app/container"
	"net/http"
)

type RestAdapter struct {
	container container.Container
}

func NewRestEngine(c container.Container) *gin.Engine {
	r := gin.Default()
	s := &RestAdapter{container: c}

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", s.pingHandler)
	}

	return r
}

func (r RestAdapter) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

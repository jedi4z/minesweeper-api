package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/minesweeper-api/app/container"
	"github.com/jedi4z/minesweeper-api/app/models"
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
		v1.POST("/games", s.createGameHandler)
	}

	return r
}

func (r RestAdapter) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (r RestAdapter) createGameHandler(c *gin.Context) {
	game := new(models.Game)

	if err := c.ShouldBindJSON(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.GameUseCases.CreateGame(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

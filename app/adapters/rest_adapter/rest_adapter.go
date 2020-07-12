package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/minesweeper-api/app/container"
	"github.com/jedi4z/minesweeper-api/app/models"
	"net/http"
	"strconv"
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
		v1.GET("/games", s.listGamesHandler)
		v1.GET("/games/:game_id", s.retrieveGameHandler)
		v1.POST("/games/:game_id/hold", s.holdGameHandler)
		v1.POST("/games/:game_id/resume", s.resumeGameHandler)
		v1.POST("/games/:game_id/flag/:cell_id", s.flagCellHandler)
		v1.POST("/games/:game_id/uncover/:cell_id", s.uncoverCellHandler)
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

func (r RestAdapter) listGamesHandler(c *gin.Context) {
	games, err := r.container.GameUseCases.ListGames()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, games)
}

func paramUint(c *gin.Context, key string) (uint, error) {
	sid := c.Param(key)
	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (r RestAdapter) retrieveGameHandler(c *gin.Context) {
	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) holdGameHandler(c *gin.Context) {
	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.HoldGame(gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) resumeGameHandler(c *gin.Context) {
	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.ResumeGame(gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) flagCellHandler(c *gin.Context) {
	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	cellID, err := paramUint(c, "cell_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.GameUseCases.FlagCell(game, cellID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) uncoverCellHandler(c *gin.Context) {
	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	cellID, err := paramUint(c, "cell_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.GameUseCases.UncoverCell(game, cellID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

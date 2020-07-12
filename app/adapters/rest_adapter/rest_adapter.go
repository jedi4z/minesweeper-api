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
		v1.POST("/users/register", s.registerUserHandler)
		v1.POST("/users/auth", s.authUserHandler)

		userResource := v1.Use(CredentialExtractorMiddleware())
		{
			userResource.POST("/games", s.createGameHandler)
			userResource.GET("/games", s.listGamesHandler)
			userResource.GET("/games/:game_id", s.retrieveGameHandler)
			userResource.POST("/games/:game_id/hold", s.holdGameHandler)
			userResource.POST("/games/:game_id/resume", s.resumeGameHandler)
			userResource.POST("/games/:game_id/flag/:cell_id", s.flagCellHandler)
			userResource.POST("/games/:game_id/uncover/:cell_id", s.uncoverCellHandler)
		}
	}

	return r
}

func (r RestAdapter) pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (r RestAdapter) checkUserCredentials(c *gin.Context) (*models.User, error) {
	accessToken := c.MustGet(AccessTokenKey).(string)
	return r.container.UserUseCases.FindByAccessToken(accessToken)
}

func (r RestAdapter) registerUserHandler(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.UserUseCases.RegisterUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (r RestAdapter) authUserHandler(c *gin.Context) {
	user := new(models.User)

	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := r.container.UserUseCases.AuthenticateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}

func (r RestAdapter) createGameHandler(c *gin.Context) {
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

	game := new(models.Game)
	if err := c.ShouldBindJSON(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game.User = user

	if err := r.container.GameUseCases.CreateGame(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) listGamesHandler(c *gin.Context) {
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

	games, err := r.container.GameUseCases.ListGames(user)
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
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(user, gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) holdGameHandler(c *gin.Context) {
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(user, gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.GameUseCases.HoldGame(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) resumeGameHandler(c *gin.Context) {
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

	gameID, err := paramUint(c, "game_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	game, err := r.container.GameUseCases.GetGame(user, gameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.container.GameUseCases.ResumeGame(game); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (r RestAdapter) flagCellHandler(c *gin.Context) {
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

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

	game, err := r.container.GameUseCases.GetGame(user, gameID)
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
	user, err := r.checkUserCredentials(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errUserUnauthorized.Error()})
		return
	}

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

	game, err := r.container.GameUseCases.GetGame(user, gameID)
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

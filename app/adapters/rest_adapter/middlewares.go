package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CredentialExtractorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(AuthorizationHeaderKey)

		if accessToken == "" {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": errCredentialRequired.Error()},
			)
			return
		}

		fs := strings.Fields(accessToken)
		if fs[0] != BearerAccessTokenKey {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": errInvalidAccessToken.Error()},
			)
			return
		}

		c.Set(AccessTokenKey, fs[1])
		c.Next()
	}
}

package rest_adapter

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func CredentialExtractorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(AuthorizationHeaderKey)

		if accessToken == "" {
			restErr := newBadRequestRestError(errCredentialRequired)
			c.AbortWithStatusJSON(restErr.StatusCode, restErr)
			return
		}

		fs := strings.Fields(accessToken)
		if fs[0] != BearerAccessTokenKey {
			restErr := newBadRequestRestError(errInvalidAccessToken)
			c.AbortWithStatusJSON(restErr.StatusCode, restErr)
			return
		}

		c.Set(AccessTokenKey, fs[1])
		c.Next()
	}
}

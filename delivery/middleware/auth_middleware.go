package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"hacktiv-assignment-final/utils/security"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		err := c.ShouldBindHeader(&header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid header " + err.Error(),
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(header.AuthorizationHeader, "Bearer ", "", 1)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token string",
			})
			c.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := security.GetIdFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "failed",
				"error":  "invalid token " + err.Error(),
			})
			c.Abort()
			return
		}

		idParam, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if id != idParam {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "failed",
				"error":  "You don't have permission to access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

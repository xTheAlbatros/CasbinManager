package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := casdoorsdk.ParseJwtToken(tokenString)
		if err != nil {
			log.Printf("Error parsing JWT token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		username := claims.Name
		if username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			c.Abort()
			return
		}

		path := c.Request.URL.Path
		method := c.Request.Method

		if ok, err := e.Enforce(username, path, method); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking permissions"})
			c.Abort()
		} else if ok {
			c.Next()
			log.Printf("Permission granted for user %s", username)
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			log.Printf("Permission denied for user %s", username)
			c.Abort()
		}
	}
}

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ExtractBearerToken maps Authorization: Bearer <token> to X-JWT header
// so that go-pkgz/auth can read the token from its expected header
func ExtractBearerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			c.Request.Header.Set("X-JWT", token)
		}
		c.Next()
	}
}

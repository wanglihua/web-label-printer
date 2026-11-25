package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
)

// CrossSite allows cross-origin requests
func CrossSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logrus.Debug("middleware cross site begin")

		c.Header("Access-Control-Allow-Origin", "*") // replace * with the specified domain
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}

		// logrus.Debug("middleware cross site end")
	}
}

package middleware

import (
	"strings"

	"printer/config"

	"github.com/gin-gonic/gin"
)

func Cache() gin.HandlerFunc {

	return func(c *gin.Context) {
		var shouldCache = false

		if strings.HasPrefix(c.Request.URL.Path, config.StaticRouteUrl) {
			shouldCache = true
		}

		if strings.HasPrefix(c.Request.RequestURI, "/favicon.") {
			shouldCache = true
		}

		if shouldCache {
			// this is placed before c.Next(); if a specific route handler sets Cache-Control,
			// the Cache-Control here will not take effect
			c.Header("Cache-Control", "public, max-age=31536000") // cache for one year
		}

		c.Next()
	}
}

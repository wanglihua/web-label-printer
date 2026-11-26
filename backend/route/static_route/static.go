package static_route

import "github.com/gin-gonic/gin"

func Route(router *gin.RouterGroup) {
	staticWwwRootRoute(router)
}

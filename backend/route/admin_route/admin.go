package admin_route

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	adminHomeDataRoute(router)
}

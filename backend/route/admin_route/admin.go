package admin_route

import (
	"printer/route/admin_route/home_route"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	home_route.Route(router)
}

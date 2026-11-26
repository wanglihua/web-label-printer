package home_route

import "github.com/gin-gonic/gin"

func Route(router *gin.RouterGroup) {
	homeDataRoute(router)
}

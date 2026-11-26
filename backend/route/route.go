package route

import (
	"printer/route/admin_route"
	"printer/route/static_route"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	var adminRouter = engine.Group("/api/admin")
	admin_route.Route(adminRouter)

	var staticRouter = engine.Group("")
	static_route.Route(staticRouter)
}

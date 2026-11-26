package home_route

import (
	"fmt"
	"net/http"

	"printer/core"
	"printer/route/action"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type HomeDataActionParam struct {
}

type HomeDataActionResult struct {
	core.ResultMessage
}

func homeDataRoute(router *gin.RouterGroup) {
	type ActionParam = HomeDataActionParam
	type ActionResult = HomeDataActionResult

	var actionConfig = &action.Config{}

	router.POST("home_data", action.Handle(actionConfig, func(c *gin.Context) error {
		var param = ActionParam{}
		var err = c.ShouldBindWith(&param, binding.JSON)
		if err != nil {
			err = fmt.Errorf("data binding error, %w", err)
			logrus.Error(err)
			return err
		}

		c.JSON(http.StatusOK, ActionResult{
			Success: true,
			Message: "success",
		})

		return nil
	}))
}

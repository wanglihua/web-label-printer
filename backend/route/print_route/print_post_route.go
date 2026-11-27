package print_route

import (
	"fmt"
	"net/http"

	"printer/core"
	"printer/route/action"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func printPostRoute(router *gin.RouterGroup) {
	type ActionParam struct{}

	type ActionResult struct {
		core.ResultMessage
	}

	var actionConfig = &action.Config{}

	router.POST("print", action.Handle(actionConfig, func(c *gin.Context) error {
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

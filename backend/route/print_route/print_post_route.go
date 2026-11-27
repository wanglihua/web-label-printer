package print_route

import (
	"fmt"
	"net/http"
	"strings"

	"printer/core"
	"printer/route/action"
	"printer/service/print_service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func printPostRoute(router *gin.RouterGroup) {
	type ActionParam struct {
		Pdf   string `form:"pdf"`
		Token string `form:"token"`
	}

	var validateActionParam = func(param *ActionParam) error {
		var pdf = strings.TrimSpace(param.Pdf)
		if pdf == "" {
			return fmt.Errorf("pdf url to print should not be empty")
		}

		var token = strings.TrimSpace(param.Token)
		if token == "" {
			return fmt.Errorf("token should not be empty")
		}

		return nil
	}

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

		err = validateActionParam(&param)
		if err != nil {
			return core.ValidationError(fmt.Sprint(err))
		}

		var pdfUrl = param.Pdf
		var token = param.Token

		err = print_service.Print(pdfUrl, token)
		if err != nil {
			return core.HandleError(err)
		}

		c.JSON(http.StatusOK, ActionResult{
			Success: true,
			Message: "success",
		})

		return nil
	}))
}

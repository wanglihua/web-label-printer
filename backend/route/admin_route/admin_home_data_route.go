package admin_route

import (
	"fmt"
	"net/http"

	"printer/core"
	"printer/core/id_gen"
	"printer/core/time_token"
	"printer/route/action"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func adminHomeDataRoute(router *gin.RouterGroup) {
	type ActionParam struct{}

	type ActionResult struct {
		core.ResultMessage
		PdfUrl string `json:"pdfUrl"`
	}

	var actionConfig = &action.Config{}

	router.POST("home_data", action.Handle(actionConfig, func(c *gin.Context) error {
		var param = ActionParam{}
		var err = c.ShouldBindWith(&param, binding.JSON)
		if err != nil {
			err = fmt.Errorf("data binding error, %w", err)
			logrus.Error(err)
			return err
		}

		var id = id_gen.NewId()
		var token = time_token.CreateToken()
		var pdfUrl = fmt.Sprintf("/api/admin/pdf/%d/%s", id, token)

		c.JSON(http.StatusOK, ActionResult{
			Success: true,
			Message: "success",
			PdfUrl:  pdfUrl,
		})

		return nil
	}))
}

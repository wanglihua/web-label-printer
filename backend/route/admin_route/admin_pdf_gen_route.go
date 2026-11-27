package admin_route

import (
	"fmt"
	"printer/core"
	"printer/route/action"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"
	"github.com/sirupsen/logrus"
)

func adminPdfGenRoute(router *gin.RouterGroup) {
	type ActionParam struct {
		Id    int64  `uri:"id"`
		Token string `uri:"token"`
	}

	var validateActionParam = func(param *ActionParam) error {
		if param.Id <=0 {
			return fmt.Errorf("id should not be empty")
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

	router.GET("pdf/:id/:token", action.Handle(actionConfig, func(c *gin.Context) error {
		var param = ActionParam{}
		var err = c.ShouldBindUri(&param)
		if err != nil {
			err = fmt.Errorf("data binding error, %w", err)
			logrus.Error(err)
			return err
		}

		err = validateActionParam(&param)
		if err != nil {
			return core.ValidationError(fmt.Sprint(err))
		}

		// todo check time token


		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
		pdf.AddPage()

		// 1. Blue rectangle (unclipped, for reference)
		pdf.SetFillColor(200, 200, 255)
		pdf.RectFromUpperLeftWithStyle(50, 50, 100, 100, "F")

		// 2. Red rectangle clipped to triangle
		pdf.SaveGraphicsState() // Save state before clipping
		pdf.ClipPolygon([]gopdf.Point{
			{X: 250, Y: 50},  // top
			{X: 350, Y: 200}, // bottom-right
			{X: 150, Y: 200}, // bottom-left
		})
		pdf.SetFillColor(255, 0, 0)
		pdf.RectFromUpperLeftWithStyle(150, 50, 200, 200, "F")
		pdf.RestoreGraphicsState() // Restore state - clipping no longer active

		// 3. Green rectangle (NOT clipped - proves RestoreGraphicsState works)
		pdf.SetFillColor(0, 200, 0)
		pdf.RectFromUpperLeftWithStyle(400, 50, 100, 100, "F")

		// pdf.WritePdf("clip-polygon.pdf")
		_, err = pdf.WriteTo(c.Writer)
		if err != nil {
			return core.HandleError(err)
		}

		return nil
	}))
}

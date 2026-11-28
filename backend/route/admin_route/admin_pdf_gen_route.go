package admin_route

import (
	"fmt"
	"math"
	"printer/core"
	"printer/route/action"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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
		if param.Id <= 0 {
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

		// Create the barcode
		qrCode, err := qr.Encode("Hello World", qr.M, qr.Auto)
		if err != nil {
			return core.HandleError(err)
		}

		// draw qrcode
		{
			pdf.SaveGraphicsState()

			var xOffset = 100
			var yOffset = 300
			var scale = 3

			const pdfColorMax = 255
			const qrcodeColorMax = 65535
			var rect = qrCode.Bounds()
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {
					var pointColor = qrCode.At(x, y)
					var r, g, b, _ = pointColor.RGBA()

					var red = uint8(math.Round(float64(r * pdfColorMax / qrcodeColorMax)))
					var green = uint8(math.Round(float64(g * pdfColorMax / qrcodeColorMax)))
					var blue = uint8(math.Round(float64(b * pdfColorMax / qrcodeColorMax)))

					pdf.SetStrokeColor(red, green, blue)
					pdf.SetFillColor(red, green, blue)
					pdf.Rectangle(
						float64(x*scale+xOffset),
						float64(y*scale+yOffset),
						float64((x+1)*scale+xOffset),
						float64((y+1)*scale+yOffset),
						"DF", 0, 0)
				}
			}

			pdf.RestoreGraphicsState()
		}

		// Scale the barcode to 200x200 pixels
		qrCode, err = barcode.Scale(qrCode, 200, 200)
		if err != nil {
			return core.HandleError(err)
		}

		// pdf.WritePdf("clip-polygon.pdf")
		_, err = pdf.WriteTo(c.Writer)
		if err != nil {
			return core.HandleError(err)
		}

		return nil
	}))
}

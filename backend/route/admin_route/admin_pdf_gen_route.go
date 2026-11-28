package admin_route

import (
	"fmt"
	"math"
	"printer/core"
	"printer/core/hash"
	"printer/core/serial_code"
	"printer/route/action"
	"strings"
	"uuid"

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
		pdf.Start(gopdf.Config{
			Unit: gopdf.UnitMM,
			PageSize: gopdf.Rect{
				W: 100,
				H: 150,
			},
		})

		pdf.AddPage()

		// draw qrcode
		{
			var data = serial_code.CreateSerialCode() + hash.SHA256(uuid.New().String())[0:8]
			var top float64 = 30
			var left float64 = 10
			var width float64 = 80

			// create the qrcode
			qrcode, err := qr.Encode(data, qr.M, qr.Auto)
			if err != nil {
				return core.HandleError(err)
			}

			var rect = qrcode.Bounds()
			if rect.Dx() <= 0 {
				return core.HandleError(fmt.Errorf("qrcode width is zero"))
			}

			var scale = width / float64(rect.Dx())

			pdf.SaveGraphicsState()

			const pdfColorMax = 255
			const qrcodeColorMax = 65535
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {
					var pointColor = qrcode.At(x, y)
					var r, g, b, _ = pointColor.RGBA()

					var red = uint8(math.Round(float64(r * pdfColorMax / qrcodeColorMax)))
					var green = uint8(math.Round(float64(g * pdfColorMax / qrcodeColorMax)))
					var blue = uint8(math.Round(float64(b * pdfColorMax / qrcodeColorMax)))

					pdf.SetStrokeColor(red, green, blue)
					pdf.SetFillColor(red, green, blue)

					pdf.Rectangle(
						float64(x)*scale+left,
						float64(y)*scale+top,
						float64(x+1)*scale+left,
						float64(y+1)*scale+top,
						"DF", 0, 0)
				}
			}

			pdf.RestoreGraphicsState()
		}

		// pdf.WritePdf("clip-polygon.pdf")
		_, err = pdf.WriteTo(c.Writer)
		if err != nil {
			return core.HandleError(err)
		}

		return nil
	}))
}

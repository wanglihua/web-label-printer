package core

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"

	"printer/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandleWrap(c *gin.Context, f func() error) {
	var err = f()

	if err != nil {
		errorDesc := fmt.Sprint(err)

		if IsValidationError(err) {
			c.JSON(http.StatusOK, ResultMessage{
				Success: false,
				Message: errorDesc,
			})
			c.Abort()
			return
		}

		// generate an error code based on the time to make it easy to find the corresponding error in the logs
		// 2006-01-02 15:04:05
		errorCode := time.Now().Format("02150405") + fmt.Sprintf("%02d", rand.Intn(999))

		// write to the system log
		logrus.Errorf("error code: %s", errorCode)
		logrus.Errorf("error message: %s", errorDesc)
		logrus.Errorf("request url: %s", c.Request.RequestURI)
		logrus.Error(string(debug.Stack()))

		if config.AppRelease() == false {
			println(errorCode)
			println(errorDesc)
			println(string(debug.Stack()))
		}

		c.JSON(http.StatusOK, ResultMessage{
			Success: false,
			Message: fmt.Sprintf("System error occurred #%s, please contact the administrator! Details: %s", errorCode, errorDesc),
		})
		c.Abort()
	}
}

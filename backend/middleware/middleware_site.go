package middleware

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"printer/config"
	"printer/core"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Site is a global middleware
// this middleware executes first among all middlewares
// handles common concerns such as exceptions and logging here
func Site() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logrus.Debug("middleware site begin")

		// the Site middleware executes first, so this defer executes last
		// it should be able to catch all errors of the site
		defer func() {
			if err := recover(); err != nil {
				var errorDesc = fmt.Sprint(err)

				if core.IsValidationError(err) {
					c.JSON(http.StatusOK, core.ResultMessage{
						Success: false,
						Message: errorDesc,
					})
					c.Abort() // must be called
					return
				}

				// generate an error code based on the time to make it easy to find the corresponding error in the logs
				// 2006-01-02 15:04:05
				errorCode := time.Now().Format("02150405") + fmt.Sprintf("%02d", rand.Intn(999))

				// write to the system log
				logrus.Errorf("error code: %s", errorCode)
				logrus.Errorf("error message: %s", errorDesc)
				logrus.Error(string(debug.Stack()))

				if config.AppRelease() == false {
					fmt.Println(errorCode)
					fmt.Println(errorDesc)
					fmt.Println(string(debug.Stack()))
				}

				// this apparently cannot be sent to the client
				c.JSON(http.StatusOK, core.ResultMessage{
					Success: false,
					Message: fmt.Sprintf("System error occurred #%s, please contact the administrator! Details: %s", errorCode, errorDesc),
				})

				c.Abort() // must be called
			}
		}()

		var timeBegin time.Time

		if len(c.Errors) > 0 {
			var errs = make([]error, 0)

			for _, eStr := range c.Errors.Errors() {
				var e = fmt.Errorf("%s", eStr)
				errs = append(errs, e)
				logrus.Error(e)
			}

			var err = errors.Join(errs...)
			var errorDesc = fmt.Sprint(err)

			// generate an error code based on the time to make it easy to find the corresponding error in the logs
			// 2006-01-02 15:04:05
			errorCode := time.Now().Format("02150405") + fmt.Sprintf("%02d", rand.Intn(999))

			// write to the system log
			logrus.Errorf("error code: %s", errorCode)
			logrus.Errorf("error message: %s", errorDesc)
			logrus.Error(string(debug.Stack()))

			if config.AppRelease() == false {
				fmt.Println(errorCode)
				fmt.Println(errorDesc)
				fmt.Println(string(debug.Stack()))
			}

			c.JSON(http.StatusOK, core.ResultMessage{
				Success: false,
				Message: fmt.Sprintf("Error occurred while processing the request #%s, please contact the administrator! Details: %s", errorCode, errorDesc),
			})

			c.Abort()

			return
		}

		if config.AppRelease() == false {
			timeBegin = time.Now()

			fmt.Print("\n\n\n\n")
			fmt.Println(strings.Repeat("-", 150))
			fmt.Printf("Begin %s \n\n", c.Request.RequestURI)
		}

		c.Next()

		if config.AppRelease() == false {
			var timeDiff = time.Since(timeBegin).Seconds()

			fmt.Printf("\nEnd   %s ", c.Request.RequestURI)
			fmt.Printf("\nTotal Time: %.3f \n", timeDiff)
			fmt.Println(strings.Repeat("=", 150))
		}

		// logrus.Debug("middleware site end")
	}
}

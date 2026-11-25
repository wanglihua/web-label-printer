package action

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime/debug"
	"time"

	"printer/config"
	"printer/core"
	"printer/core/auth_type"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Transaction bool
	AuthType    auth_type.AuthType
}

func Handle(cfg *Config, fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
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

				c.JSON(http.StatusOK, core.ResultMessage{
					Success: false,
					Message: fmt.Sprintf("System error occurred #%s, please contact the administrator! Details: %s", errorCode, errorDesc),
				})

				c.Abort() // must be called
			}
		}()

		var transaction = false
		var at = auth_type.NoLoginRequired

		if cfg != nil {
			transaction = cfg.Transaction
			at = cfg.AuthType
		}

		core.ErrorHandleWrap(c, func() error {
			// maybe check auth type
			_ = at

			// maybe start a database transaction
			_ = transaction

			return fn(c)
		})
	}
}

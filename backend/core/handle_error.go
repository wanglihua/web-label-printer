package core

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func HandleError(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		logrus.Error(fmt.Sprintf("%s#%d", file, line))
	}

	logrus.Error(err)
	logrus.Error(string(debug.Stack()))

	return err
}

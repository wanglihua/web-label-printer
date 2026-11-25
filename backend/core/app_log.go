package core

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"sort"
	"time"

	"printer/config"

	file_rotate_logs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// LogAppStartupError records errors during application initialization, not during HTTP requests. The two are recorded somewhat differently.
func LogAppStartupError(err error) {
	logrus.Error(err)
	logrus.Error(string(debug.Stack()))
	fmt.Println(fmt.Sprint(err))
	fmt.Println(string(debug.Stack()))

	os.Exit(1) // this exit still feels necessary
}

func SetupAppLog() {
	logsPath := getLogsPath()
	if err := os.MkdirAll(logsPath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(logsPath)

	rotateLogFile, err := file_rotate_logs.New(
		filepath.Join(logsPath, "%Y_%m_%d.log"),
		// fileRotateLogs.WithLinkName("/var/log/cli.log"), // create a symlink pointing to the latest log file
		// MaxAge and RotationCount cannot be both set
		// fileRotateLogs.WithMaxAge(6*time.Minute), // clear, with a minimum unit of minutes
		file_rotate_logs.WithRotationCount(100),        // number defaults to 100 files; cleanup starts once there are more than 100 or the cleanup time is reached
		file_rotate_logs.WithRotationTime(time.Minute), // rotate polls at a minimum of 1 minute; defaults to 60s, values below 1 minute are treated as 1 minute
	)

	if err != nil {
		log.Printf("failed to create fileRotateLogs: %s", err)
		return
	}

	logrus.SetReportCaller(true)
	logrus.SetOutput(rotateLogFile)

	// set the log level
	if config.AppRelease() {
		logrus.SetLevel(logrus.ErrorLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		ForceColors:     false,
		DisableQuote:    true,
		DisableSorting:  false,
		SortingFunc: func(s []string) {
			priority := map[string]int{
				"time":  1,
				"level": 2,
				"file":  3,
				"func":  4,
				"msg":   5,
			}
			sort.SliceStable(s, func(i, j int) bool {
				pi, pj := priority[s[i]], priority[s[j]]
				switch {
				case pi > 0 && pj > 0:
					return pi < pj
				case pi > 0:
					return true
				case pj > 0:
					return false
				default:
					return s[i] < s[j]
				}
			})
		},
	})
}

func getLogsPath() string {
	return path.Join(GetAppPath(), "logs")
}

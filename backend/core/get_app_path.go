package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var appPath string = ""

func GetAppPath() string {
	if appPath == "" {
		exeFile, err := os.Executable()
		if err != nil {
			fmt.Println(fmt.Sprint(err))
			logrus.Fatal(err)
		}

		exePath := filepath.Dir(exeFile)

		// this is another way to get exePath
		// exePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		// if err != nil {
		// 	fmt.Println(fmt.Sprint(err))
		// 	logrus.Fatal(err)
		// }

		// Is Symlink
		fi, err := os.Lstat(exePath)
		if err != nil {
			fmt.Println(fmt.Sprint(err))
			logrus.Fatal(err)
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			exePath, err = os.Readlink(exePath)
			if err != nil {
				fmt.Println(fmt.Sprint(err))
				logrus.Fatal(err)
			}
		}

		// Is Current Directory
		if exePath == "." {
			exePath, err = os.Getwd()
			if err != nil {
				fmt.Println(fmt.Sprint(err))
				logrus.Fatal(err)
			}
		}

		// exePath = strings.ReplaceAll(exePath, "\\", "/")

		appPath = exePath
	}

	return appPath
}

package config

const AppModeEnv = "APP_MODE"

var appMode = ""

// AppMode uses the APP_MODE environment variable if it is set, otherwise defaults to release.
func AppMode() string {
	if appMode == "" {
		appMode = GetEnv(AppModeEnv, "release")
	}

	return appMode
}

func AppRelease() bool {
	if AppMode() == "release" {
		return true
	}

	return false
}

func AppDebug() bool {
	if AppMode() == "debug" {
		return true
	}

	return false
}

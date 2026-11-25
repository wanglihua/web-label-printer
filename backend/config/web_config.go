package config

const SessionName = "session_id"
const SessionSecret = "6f3b8ded65bd7a4db11625ac84e579bb"
const SessionTimeoutSeconds = 60 * 60 * 24
const StaticRouteUrl = "/static"

const httpPortDefault = "8113"
const httpsPortDefault = "8114"
const httpPortEnv = "WEB_LABEL_PRINTER_HTTP_PORT"
const httpsPortEnv = "WEB_LABEL_PRINTER_HTTPS_PORT"

var httpPort = ""

func GetHttpPort() string {
	if httpPort == "" {
		httpPort = GetEnv(httpPortEnv, httpPortDefault)
	}

	return httpPort
}

var httpsPort = ""

func GetHttpsPort() string {
	if httpsPort == "" {
		httpsPort = GetEnv(httpsPortEnv, httpsPortDefault)
	}

	return httpsPort
}

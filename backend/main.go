package main

import (
	"os"

	"printer/config"
	"printer/core"
	"printer/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.LoadEnvFile()
	if err != nil {
		core.LogAppStartupError(err)
		os.Exit(1)
	}

	core.SetupAppLog()

	if config.AppDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var engine = gin.Default()

	var store sessions.Store

	store = memstore.NewStore([]byte(config.SessionSecret))
	store.Options(sessions.Options{
		Path:   "/", // this must be set, otherwise different session ids may be generated under different paths
		MaxAge: config.SessionTimeoutSeconds,
	})

	engine.Use(sessions.Sessions(config.SessionName, store))

	// place these middlewares after the session configuration above so that session can be used in the middlewares
	// the topmost middleware executes first and exits last

	// this must be placed after engine.Use(middleware.Site()),
	// otherwise, after an exception is caught in Site(), the compression process will be interrupted unexpectedly, and the client will not receive the response data
	// engine.Use(gzip.Gzip(gzip.DefaultCompression))

	engine.Use(middleware.Site())
	engine.Use(middleware.CrossSite())

	// if configured as a backend behind a reverse proxy, this gzip compression is not needed
	engine.Use(gzip.Gzip(gzip.DefaultCompression)) // keep it here, the main errors are handled by action.Handle()
}

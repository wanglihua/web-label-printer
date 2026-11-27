package static_route

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"printer/config"
	"printer/core"
	"printer/route/action"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*

Pattern Matching Rules
The embed directive supports several pattern types:

file.txt - exact file match
dir/* - all entries directly in a directory; if a matched entry is a directory, its non-hidden subtree is embedded
dir/** - not supported; use dir instead
dir - the directory and all contents recursively

*/

//go:embed www_root
var wwwRootFs embed.FS

func staticWwwRootRoute(router *gin.RouterGroup) {
	staticFs, err := fs.Sub(wwwRootFs, "www_root/static")
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		logrus.Fatal(err)
	}

	router.StaticFS(config.StaticRouteUrl, http.FS(staticFs))

	var actionConfig = &action.Config{}

	router.GET("/favicon.webp", action.Handle(actionConfig, func(c *gin.Context) error {
		fileContent, err := wwwRootFs.ReadFile("www_root/favicon.webp")
		if err != nil {
			return core.HandleError(err)
		}

		c.Data(http.StatusOK, "image/webp", fileContent)

		return nil
	}))

	/*
		HTTP redirects
		Use c.Redirect with an appropriate HTTP status code to redirect the client:

		301 (http.StatusMovedPermanently)  — the resource has permanently moved. Browsers and search engines update their caches.
		302 (http.StatusFound)             — temporary redirect. The browser follows but does not cache the new URL.
		307 (http.StatusTemporaryRedirect) — like 302, but the browser must preserve the original HTTP method (useful for POST redirects).
		308 (http.StatusPermanentRedirect) — like 301, but the browser must preserve the original HTTP method.
	*/

	var indexHtmlRoute = action.Handle(actionConfig, func(c *gin.Context) error {
		file, err := wwwRootFs.ReadFile("www_root/index.html")
		if err != nil {
			logrus.Error(err)
			return err
		}

		// the index.html should not be cached
		c.Header("Cache-Control", "public, max-age=0")
		c.Data(http.StatusOK, "text/html", file)

		return nil
	})

	router.GET("/", indexHtmlRoute)
	router.GET("/admin", indexHtmlRoute)
	router.GET("/admin/*any", indexHtmlRoute)
}

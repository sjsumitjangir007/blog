package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func ConfigRoute(router *gin.Engine) {
	router.Use(apmgin.Middleware(router))
	router.HTMLRender = ginview.Default()
	router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.StaticFile("/manifest.json", "./resources/manifest.json")
	router.StaticFile("/sw.js", "./resources/sw.js")
	router.LoadHTMLGlob("templates/**/*.html")
}

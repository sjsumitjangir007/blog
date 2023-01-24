package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Home struct {
	Path string
}

func (home *Home) Handler(c *gin.Context) {
	c.Header("Cache-Control:", "max-age=31536000")
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "Welcome",
	})
}

var homeRoute Home = Home{Path: "/"}

func HomeRoute(router *gin.Engine) {
	router.GET(homeRoute.Path, homeRoute.Handler)
}

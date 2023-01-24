package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Editor struct {
	Path string
}

func (editor *Editor) Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "editors/index.html", gin.H{
		"title": "Welcome Editor",
	})
}

var editorRoute Editor = Editor{Path: "/editor"}

func EditorRoute(router *gin.Engine) {
	router.GET(editorRoute.Path, editorRoute.Handler)
}

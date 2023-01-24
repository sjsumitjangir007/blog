package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	Path string
}

func (admin *Admin) Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "Welcome Admin",
	})
}

var adminRoute Admin = Admin{Path: "/admin"}

func AdminRoute(router *gin.Engine) {
	router.GET(adminRoute.Path, adminRoute.Handler)
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Path string
}

func (user *User) Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"title": "Welcome User",
	})
}

var userRoute User = User{Path: "/users"}

func UserRoute(router *gin.Engine) {
	router.GET(userRoute.Path, userRoute.Handler)
}

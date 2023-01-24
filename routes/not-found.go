package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundRoute(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "not-found/index.html", gin.H{"title": "Page Not Found", "message": "Page not found"})
	})
}

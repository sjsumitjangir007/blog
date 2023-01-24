package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Path string
}

func (admin *Contact) Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "contact/index.html", gin.H{
		"title": "Contact Us",
	})
}

var contactRoute Contact = Contact{Path: "/contact"}

func ContactRoute(router *gin.Engine) {
	router.GET(contactRoute.Path, contactRoute.Handler)
}

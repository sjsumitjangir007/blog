package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Welcome",
		})
	})
	router.GET("/posts/post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/post.tmpl", gin.H{
			"title": "Single Post",
		})
	})
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "not-found/index.tmpl", gin.H{"title": "Page Not Found", "message": "Page not found"})
	})
	router.Run(":8080")
}

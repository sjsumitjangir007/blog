package main

import (
	"github.com/gin-gonic/gin"

	blogRoutes "github.com/sjsumitjangir007/blog/routes"
)

func main() {
	router := gin.Default()

	blogRoutes.ConfigRoute(router)

	blogRoutes.HomeRoute(router)
	blogRoutes.PostRoute(router)
	blogRoutes.PostsRoute(router)
	blogRoutes.AdminRoute(router)
	blogRoutes.EditorRoute(router)
	blogRoutes.UserRoute(router)
	blogRoutes.ContactRoute(router)
	blogRoutes.NotFoundRoute(router)

	router.Run(":8081")
}

package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sjsumitjangir007/blog/util"
)

type Posts struct {
	Path string
}

func (post *Posts) Handler(c *gin.Context) {
	postsContent, err := getLocalBlogDB()
	if err != nil {
		log.Fatal("red dir issue: ", err)
		c.Error(err)
		c.HTML(http.StatusNotFound, "not-found/index.html", gin.H{"title": "Page Not Found", "message": "Unable to provide page"})
	} else {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
			"posts": postsContent,
		})
	}
}

var postsRoute Posts = Posts{Path: "/posts"}

func PostsRoute(router *gin.Engine) {
	router.GET(postsRoute.Path, postsRoute.Handler)
}

func getLocalBlogDB() (*[]util.LocalMdMetadata, error) {
	// files, err := os.ReadDir("./local_db/blogs")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dirNames := []string{}
	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// 	dirNames = append(dirNames, f.Name())
	// }
	// return dirNames, err

	content, err := os.ReadFile("./local_db/blogs/metadata.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}

	// Now let's unmarshall the data into `payload`
	var payload []util.LocalMdMetadata
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}

	return &payload, nil
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	htmlTemplate "html/template"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/sjsumitjangir007/blog/util"
)

type Post struct {
	Path string
}

func (post *Post) Handler(c *gin.Context) {
	name := c.Param("name")
	metadata, err := getLocalBlogMetadata("local_db/blogs/" + name)
	if err != nil {
		fmt.Println("metadata no found")
		// log.Fatal("metadata Error: ", err)
		// c.Error(err)
		c.HTML(http.StatusNotFound, "not-found/index.html", gin.H{"title": "Page Not Found", "message": "Page description not found"})
	} else {
		html, errr := getLocalBlogContent("local_db/blogs/" + name + "/" + util.OptStr(metadata.ContentFilePath, "index"))
		if errr != nil {
			fmt.Println("inside error")
			// log.Fatal("html Error: ", err)
			// c.Error(err)
			c.HTML(http.StatusNotFound, "not-found/index.html", gin.H{"title": "Page Not Found", "message": "Page not found"})
		} else {
			c.HTML(http.StatusOK, "posts/post.html", gin.H{
				"title":    "Single Post",
				"Content":  htmlTemplate.HTML(html),
				"metadata": metadata,
				"tags":     strings.Split(metadata.Tags, ", "),
			})
		}
	}
}

var postRoute Post = Post{Path: "/post/:name"}

func PostRoute(router *gin.Engine) {
	router.GET(postRoute.Path, postRoute.Handler)
}

func getLocalBlogMetadata(path string) (*util.LocalMdMetadata, error) {
	content, errr := os.ReadFile(path + "/metadata.json")
	if errr != nil {
		// log.Fatal("Error when opening file: ", err)
		return nil, errr
	} else {
		// Now let's unmarshall the data into `payload`
		var payload util.LocalMdMetadata
		errr = json.Unmarshal(content, &payload)
		if errr != nil {
			// log.Fatal("Error during Unmarshal(): ", err)
			return nil, errr
		} else {
			return &payload, nil
		}
	}
}

func getLocalBlogContent(path string) (string, error) {
	mdByteArr, err := os.ReadFile(path + ".md") // just pass the file name
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	output := markdown.ToHTML(mdByteArr, nil, nil)

	return string(output), nil
}

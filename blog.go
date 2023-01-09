package main

import (
	"encoding/json"
	"fmt"
	htmlTemplate "html/template"
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

type LocalMdMetadata struct {
	ModifiedOn         string `json:"modifiedOn"`
	CreatedOn          string `json:"createdOn"`
	AuthorName         string `json:"authorName"`
	AuthorEmail        string `json:"authorEmail"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Tags               string `json:"tags"`
	OgType             string `json:"ogType"`
	OgURL              string `json:"ogURL"`
	OgDescription      string `json:"ogDescription"`
	OgImage            string `json:"ogImage"`
	TwitterCard        string `json:"twitterCard"`
	TwitterDescription string `json:"twitterDescription"`
	TwitterTitle       string `json:"twitterTitle"`
	TwitterURL         string `json:"twitterURL"`
	TwitterIamge       string `json:"twitterIamge"`
}

func main() {

	router := gin.Default()
	router.HTMLRender = ginview.Default()
	router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Welcome",
		})
	})
	router.GET("/post/:name", func(c *gin.Context) {
		name := c.Param("name")
		html, err := getLocalBlogContent("local_db/" + name)
		if err != nil {
			log.Fatal("html Error: ", err)
			c.Error(err)
			return
		}
		metadata, err := getLocalBlogMetadata("local_db/" + name)
		if err != nil {
			log.Fatal("metadata Error: ", err)
			c.Error(err)
			return
		}
		fmt.Println(metadata.Tags)
		c.HTML(http.StatusOK, "posts/post.tmpl", gin.H{
			"title":    "Single Post",
			"Content":  htmlTemplate.HTML(html),
			"metadata": metadata,
		})

		// c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})
	router.GET("/posts", func(c *gin.Context) {
		dirs, err := getLocalBlogDB()
		if err != nil {
			log.Fatal("red dir issue: ", err)
			c.Error(err)
			return
		}
		fmt.Println(dirs)
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title":     "Posts",
			"postsName": dirs,
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
	router.Run(":8081")
}

func getLocalBlogMetadata(path string) (*LocalMdMetadata, error) {
	content, err := os.ReadFile(path + "/metadata.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}

	// Now let's unmarshall the data into `payload`
	var payload LocalMdMetadata
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}

	return &payload, nil
}

func getLocalBlogContent(path string) (string, error) {
	mdByteArr, err := os.ReadFile(path + "/index.md") // just pass the file name
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	output := markdown.ToHTML(mdByteArr, nil, nil)

	return string(output), nil
}

func getLocalBlogDB() ([]string, error) {
	files, err := os.ReadDir("./local_db")
	if err != nil {
		log.Fatal(err)
	}

	dirNames := []string{}
	for _, f := range files {
		fmt.Println(f.Name())
		dirNames = append(dirNames, f.Name())
	}
	return dirNames, err
}

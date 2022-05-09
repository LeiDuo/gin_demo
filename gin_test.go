package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	//默认路由引擎
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
			"title":   "hi!",
		})
	})
	_ = engine.Run(":9090")
}

func TestGinMethod(t *testing.T) {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	_ = r.Run(":9091")
}

func TestTemplateHTML(t *testing.T) {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "你请求了posts/index.html",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "你请求了users/index.html",
		})
	})
	_ = r.Run(":8080")
}

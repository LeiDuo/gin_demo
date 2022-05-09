package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"html/template"
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

func TestNotEscapedTemplate(t *testing.T) {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("templates/index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://baidu.com'>雷子卓的博客</a>")
	})

	_ = router.Run(":8080")
}

func TestStaticTemplate(t *testing.T) {
	router := gin.Default()
	router.Static("/test", "./static")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/image", func(c *gin.Context) {
		c.HTML(http.StatusOK, "image.html", gin.H{
			"title": "你请求了image.html",
		})
	})
	_ = router.Run(":9090")
}
func TestDataRender(t *testing.T) {
	router := gin.Default()
	//JSON渲染
	router.GET("/simple_json", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	router.GET("/struct_json", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})

	//XML渲染
	router.GET("/simple_xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "json"})
	})
	router.GET("/struct_xml", func(c *gin.Context) {
		// 方法二：使用结构体,但是不能使用匿名结构体,必须先声明类型后使用
		type msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		var msgXml msg
		msgXml.Name = "小王子"
		msgXml.Message = "Hello world!"
		msgXml.Age = 18
		c.XML(http.StatusOK, msgXml)
	})

	// YAML渲染
	router.GET("/simple_yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK, "title": "title"})
	})

	// protobuf渲染
	router.GET("/simple_protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	_ = router.Run(":9090")
}

func TestParam(t *testing.T) {
	router := gin.Default()
	// url参数
	router.GET("/get", func(c *gin.Context) {
		// 可以添加默认值
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// form参数
	router.POST("/post", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// path value参数
	router.DELETE("/delete/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	_ = router.Run(":9090")
}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func TestBind(t *testing.T) {
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	router.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	_ = router.Run(":9090")
}

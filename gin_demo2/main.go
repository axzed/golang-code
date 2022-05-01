package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "./static/static")
	// gin框架给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 模板解析
	// r.LoadHTMLFiles("template/posts/index.tmpl", "template/users/index.tmpl")
	// 正则全部加载
	r.LoadHTMLGlob("template/**/*")
	r.GET("/posts/index", func(ctx *gin.Context) {
		// HTTP请求
		// gin.H 就是一个 map[string]interface{} 很方便
		ctx.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "posts/index",
		})
	})
	r.GET("/users/index", func(ctx *gin.Context) {
		// HTTP请求
		// gin.H 就是一个 map[string]interface{} 很方便
		ctx.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href = 'http://axzed.com'>xuewenchao</a>",
		})
	})

	r.GET("/home", func(ctx *gin.Context){
		ctx.HTML(http.StatusOK, "home.tmpl", nil)
	})
	r.Run(":5050") // 启动server
}

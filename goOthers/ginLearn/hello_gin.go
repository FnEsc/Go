package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})
	r.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		fmt.Println(action)
		//  截取/
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	r.GET("user", func(context *gin.Context) {
		//指定默认值 //http://127.0.0.1/user
		//或者这样访问 http://localhost:8080/user?name=TA
		name := context.DefaultQuery("name", "you")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	//3.监听端口，默认8080
	r.Run()
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	/*engine.GET("/hello", func(ctx *gin.Context) {
		fmt.Println("请求路径:", ctx.FullPath())

		ctx.Writer.Write([]byte("Hello,gin\n"))
	})*/

	engine.Handle("GET", "/hello", func(ctx *gin.Context) {
		path := ctx.FullPath()
		fmt.Printf("path: %v\n", path)
		name := ctx.DefaultQuery("name", "Tom")
		fmt.Println(name)

		ctx.Writer.Write([]byte("success," + name))
	})

	engine.Handle("POST", "/login", func(ctx *gin.Context) {
		path := ctx.FullPath()
		fmt.Printf("path: %v\n", path)
		username := ctx.PostForm("username")
		passwd := ctx.PostForm("password")
		fmt.Println(username, passwd)

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "登录成功",
		})
	})

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}

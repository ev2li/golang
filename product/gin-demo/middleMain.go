package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	//方法一
	// e.Use(RequestInfos())
	//方法二
	e.GET("/query", RequestInfos(), func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("success"))
	})

	e.Run()
}

/// 打印请求信息的中间件
func RequestInfos() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method
		fmt.Println(path, method)
		ctx.Next()
		fmt.Println(ctx.Writer.Status())
	}
}

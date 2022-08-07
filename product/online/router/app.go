package router

import (
	"online/service"

	_ "online/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", service.Ping)
	//问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	//用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)

	//用户排行榜
	r.GET("/rank-list", service.GetRankList)
	//提交记录
	r.GET("/submit-list", service.GetSubmitList)

	//管理员私有方法
	r.POST("/problem-create", service.GetProblemCreate)
	return r
}

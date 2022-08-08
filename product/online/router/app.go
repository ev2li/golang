package router

import (
	"online/middlewares"
	"online/service"

	_ "online/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// Swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", service.Ping)

	// 路由规则

	// 公有方法
	// 问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)

	// 用户排行榜
	r.GET("/rank-list", service.GetRankList)
	//提交记录
	r.GET("/submit-list", service.GetSubmitList)

	// 管理员私有方法
	// 问题创建
	r.POST("/problem-create", middlewares.AuthAdminCheck(), service.GetProblemCreate)
	// 分类列表
	r.GET("/category-list", middlewares.AuthAdminCheck(), service.GetCategoryList)
	// 分类创建
	r.POST("/category-create", middlewares.AuthAdminCheck(), service.CategoryCreate)
	// 分类修改
	r.PUT("/category-modify", middlewares.AuthAdminCheck(), service.CategoryModify)
	// 分类删除
	r.DELETE("/category-delete", middlewares.AuthAdminCheck(), service.CategoryDelete)
	return r
}

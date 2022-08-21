package handler

import (
	"context"
	"gateway/internal/service"
	"gateway/pkg/e"
	"gateway/pkg/res"
	"gateway/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister用户注册
func UserRegister(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	//gin.Key中获取服务实例
	userService := ctx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

// UserLogin用户登录
func UserLogin(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	//gin.Key中获取服务实例
	userService := ctx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, _ := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Data: res.TokenData{
			User:  userResp.UserDetail,
			Token: token,
		},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

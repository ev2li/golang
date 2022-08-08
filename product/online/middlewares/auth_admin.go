package middlewares

import (
	"net/http"
	"online/helper"

	"github.com/gin-gonic/gin"
)

//AuthAdminCheck
//AuthAdmin is a middleware that checks for admin authentication
func AuthAdminCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Check if user is admin
		auth := ctx.GetHeader("Authorization")
		userClaim, err := helper.ParserToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}

		if userClaim.IsAdmin != 1 {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		ctx.Next()
	}
}

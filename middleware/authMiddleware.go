package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/global"
	"gvb_blog/models"
	"gvb_blog/response"
	"gvb_blog/utils/jwt"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// 没有直接return
		if tokenString == "" || strings.HasSuffix(tokenString, "Bearer ") {
			response.Fail(ctx, "请先登录")
			ctx.Abort()
			return
		}

		// 判断token是否有效
		tokenString = tokenString[0:7]
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Fail(ctx, "token已过期")
			ctx.Abort()
			return
		}
		var userMo models.UserModel
		err = global.DB.Take(&userMo, claims.UserId).Error
		if err != nil {
			response.Fail(ctx, "获取token用户失败")
			ctx.Abort()
			return
		}
		ctx.Set("user", claims)
		ctx.Next()
	}
}

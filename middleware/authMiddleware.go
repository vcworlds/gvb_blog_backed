package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
		tokenString = tokenString[7:]
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			global.Log.Error(err)
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
		// 判断token是否存在redis
		tokenKey := fmt.Sprintf("token_%d", claims.UserId)
		tokenR, err := global.Redis.Get(context.Background(), tokenKey).Result()
		if err == redis.Nil {
			global.Log.Error(err)
			response.Fail(ctx, "token已失效")
			ctx.Abort()
			return
		} else if err != nil {
			global.Log.Error(err)
			response.Fail(ctx, "token不存在")
			ctx.Abort()
			return
		}
		if tokenR != token.Raw {
			response.Fail(ctx, "token不正确")
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// 没有直接return
		if tokenString == "" || strings.HasSuffix(tokenString, "Bearer ") {
			response.Fail(ctx, "请先登录")
			ctx.Abort()
			return
		}

		// 判断token是否有效
		tokenString = tokenString[7:]
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			global.Log.Error(err)
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
		if claims.Role != "管理员" {
			response.Fail(ctx, "您没有该权限！如有问题请联系作者！")
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

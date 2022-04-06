package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"youyin/common"
	"youyin/response"
)

// JwtAuthMiddleware 登录验证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 如果不是登录请求，直接通过
		if c.Request.URL.Path == "/login" {
			c.Next()
			return
		}
		// 获取token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.Error(c, http.StatusUnauthorized, "请先登录")
			c.Abort()
			return
		}

		// 校验token
		_, err := common.CheckToken(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "请先登录")
			c.Abort()
			return
		}
		// 通过，放行
		c.Next()
	}
}

//func JwtAuthMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			response.Success(c, gin.H{"errno": reqstatus.LOGINERR, "errmsg": "未获取登录状态"})
//			c.Abort()
//			return
//		}
//		reqToken, err := common.ParseToken(authHeader)
//		if err != nil {
//			response.Success(c, gin.H{"errno": reqstatus.LOGINERR, "errmsg": "未获取登录状态"})
//			c.Abort()
//			return
//		}
//		c.Set("username", reqToken.Username)
//		c.Set("userId", reqToken.UserId)
//		c.Next()
//
//	}
//}

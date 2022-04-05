package middleware

import (
	"github.com/gin-gonic/gin"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/response"
)

// JwtAuthMiddleware 登录验证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Success(c, gin.H{"errno": reqstatus.LOGINERR, "errmsg": "未获取登录状态"})
			c.Abort()
			return
		}
		reqToken, err := common.ParseToken(authHeader)
		if err != nil {
			response.Success(c, gin.H{"errno": reqstatus.LOGINERR, "errmsg": "未获取登录状态"})
			c.Abort()
			return
		}
		c.Set("username", reqToken.Username)
		c.Set("userId", reqToken.UserId)
		c.Next()

	}
}

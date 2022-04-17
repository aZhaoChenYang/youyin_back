package response

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Error(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func DbError(c *gin.Context, err error) {
	zap.L().Error("db error", zap.Error(err))
	Error(c, 500, "数据库错误")
}

func ParamError(c *gin.Context, err error) {
	zap.L().Error("param error", zap.Error(err))
	Error(c, 400, "参数错误")
}

func ThirdPartyError(c *gin.Context, err error) {
	zap.L().Error("third party error", zap.Error(err))
	Error(c, 500, "第三方错误")
}

func GenTokenError(c *gin.Context, err error) {
	zap.L().Error("gen token error", zap.Error(err))
	Error(c, 500, "生成token错误")
}

func EncryptionError(c *gin.Context, err error) {
	zap.L().Error("encryption error", zap.Error(err))
	Error(c, 500, "加密错误")
}

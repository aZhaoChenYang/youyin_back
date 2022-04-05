package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装返回的代码
func Response(c *gin.Context, httpStatus int, data gin.H) {
	c.JSON(httpStatus, data)
}

// Success 成功时返回的代码
func Success(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, data)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

func AddShop(c *gin.Context) {
	var request model.Shop
	if err := c.BindJSON(&request); err != nil {
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})

}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	var tag model.Tag
	if err := c.BindJSON(&tag); err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := tag.Add(); err != nil {
		zap.L().Error("添加标签失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	var tag model.Tag
	tag.ID = uint(id)
	if err := tag.Delete(); err != nil {
		zap.L().Error("删除标签失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	var tag model.Tag
	list, err := tag.GetList()
	if err != nil {
		zap.L().Error("获取标签列表失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

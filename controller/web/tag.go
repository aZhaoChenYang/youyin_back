package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"youyin/model"
	"youyin/response"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	var tag model.Tag
	if err := c.BindJSON(&tag); err != nil {
		response.ParamError(c, err)
		return
	}
	if err := tag.Add(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.ParamError(c, err)
		return
	}
	if err := (&model.Tag{ID: uint(id)}).Delete(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// GetTagList 获取标签列表
func GetTagList(c *gin.Context) {
	list, err := (&model.Tag{}).GetList()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, list)
}

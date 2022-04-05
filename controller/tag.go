package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

type tag struct {
	Name string `bind:"required"`
}

func AddTag(c *gin.Context) {

	var request tag
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	db := common.GetDB()
	if err = db.Create(&model.Tag{Name: request.Name}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "添加失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})

}

func GetTag(c *gin.Context) {
	var tags []model.Tag
	db := common.GetDB()
	if err := db.Find(&tags).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "数据库查询失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "查询成功", "data": tags})
}

func DeleteTag(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	db := common.GetDB()
	if err := db.Where("id=?", id).Delete(&model.Tag{}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "删除失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "删除成功"})

}

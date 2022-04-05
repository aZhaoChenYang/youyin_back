package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

type people struct {
	Number int `bind:"required"`
}

func AddPeople(c *gin.Context) {
	var request people
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	db := common.GetDB()
	if err = db.Create(&model.People{Number: request.Number}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "添加失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})
}

func GetPeople(c *gin.Context) {
	var peoples []model.People
	db := common.GetDB()
	if err := db.Find(&peoples).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "数据库查询失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "查询成功", "data": peoples})
}

func DeletePeople(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	db := common.GetDB()
	if err := db.Where("id=?", id).Delete(&model.People{}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "删除失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "删除成功"})
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

func AddShop(c *gin.Context) {
	var request model.Shop
	if err := c.BindJSON(&request); err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	db := common.GetDB()
	if err := db.Create(&request).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "添加失败"})
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})

}

func GetShop(c *gin.Context) {
	var shops []model.Shop
	db := common.GetDB()
	if err := db.Find(&shops).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "查询失败"})
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "查询成功", "data": shops})

}

func UpdateShop(c *gin.Context) {
	var request model.Shop
	if err := c.BindJSON(&request); err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	db := common.GetDB()
	if err := db.Model(&model.Shop{}).Where("id=?", request.ID).Update(&request).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "修改失败"})
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "修改成功"})

}

func DeleteShop(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	db := common.GetDB()
	var shop model.Shop
	if err := db.Where("id=?", id).First(&shop).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "删除失败"})
	}
	if db.Model(&shop).Association("Swipers").Count() != 0 {
		response.Success(c, gin.H{"errno": reqstatus.DATAERR, "errmsg": "请删除关联的轮播图"})
		return
	}
	if err := db.Delete(&shop).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "删除失败"})
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "删除成功"})

}

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

func AddShop(c *gin.Context) {
	var shop model.Shop
	if err := c.BindJSON(&shop); err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	err := shop.Create()
	if err != nil {
		zap.L().Error("添加失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "添加失败")
		return
	}
	response.Success(c, "添加成功")
}

func GetShopList(c *gin.Context) {
	var shop model.Shop
	list, err := shop.GetList()
	if err != nil {
		zap.L().Error("获取失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取失败")
		return
	}
	response.Success(c, list)
}

func UpdateShop(c *gin.Context) {
	var shop model.Shop
	if err := c.BindJSON(&shop); err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	err := shop.Update()
	if err != nil {
		zap.L().Error("更新失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "更新失败")
		return
	}
	response.Success(c, "更新成功")
}

func DeleteShop(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	shop := model.Shop{
		Model: gorm.Model{ID: uint(id)},
	}
	err = shop.Delete()
	if err != nil {
		zap.L().Error("删除失败", zap.Error(err))
		if err == model.ForeignkeyError {
			response.Error(c, http.StatusBadRequest, "删除失败，请先删除该店铺下的轮播图")
		} else {
			response.Error(c, http.StatusInternalServerError, "删除失败")
		}
		return
	}
	response.Success(c, "删除成功")
}

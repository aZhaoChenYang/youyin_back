package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

//添加轮播图
func AddSwiper(c *gin.Context) {
	var swiper model.Swiper
	err := c.BindJSON(&swiper)
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = swiper.Add(); err != nil {
		zap.L().Error("添加人数失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)

}

//修改轮播图
func UpdateSwiper(c *gin.Context) {
	var swiper model.Swiper
	err := c.BindJSON(&swiper)
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = swiper.Update(); err != nil {
		zap.L().Error("修改轮播图失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// 删除轮播图
func DeleteSwiper(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	var swiper model.Swiper
	swiper.ID = uint(id)
	if err = swiper.Delete(); err != nil {
		zap.L().Error("删除轮播图失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// 获取轮播图列表
func GetSwiperList(c *gin.Context) {
	var swiper model.Swiper
	list, err := swiper.GetList()
	if err != nil {
		zap.L().Error("获取轮播图列表失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

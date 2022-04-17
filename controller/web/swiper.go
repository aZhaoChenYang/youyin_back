package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"youyin/model"
	"youyin/response"
)

//添加轮播图
func AddSwiper(c *gin.Context) {
	var swiper model.Swiper
	err := c.BindJSON(&swiper)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	if err = swiper.Add(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)

}

//修改轮播图
func UpdateSwiper(c *gin.Context) {
	var swiper model.Swiper
	err := c.BindJSON(&swiper)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	if err = swiper.Update(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// 删除轮播图
func DeleteSwiper(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.ParamError(c, err)
		return
	}
	if err = (&model.Swiper{ID: uint(id)}).Delete(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// 获取轮播图列表
func GetSwiperList(c *gin.Context) {
	list, err := (&model.Swiper{}).GetList()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, list)
}

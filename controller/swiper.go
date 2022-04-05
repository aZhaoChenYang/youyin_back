package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

func AddSwiper(c *gin.Context) {
	type reqswiper struct {
		Imgurl   string `json:"imgurl" bind:"required"`
		ShopName string `json:"shop_name" bind:"required"`
	}
	var request reqswiper
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	db := common.GetDB()
	if err := db.Create(&model.Swiper{Imgurl: request.Imgurl, ShopName: request.ShopName}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "添加失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})
}

func GetSwiper(c *gin.Context) {
	var reqswipers []model.Swiper
	db := common.GetDB()
	if err := db.Find(&reqswipers).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "查询失败"})
		return
	}
	type Swipers struct {
		ID        uint   `json:"id"`
		Img_url   string `json:"img_url"`
		Shop_name string `json:"shop_name"`
	}
	var swipers []Swipers
	for _, reqswiper := range reqswipers {

		swiper := Swipers{
			ID:        reqswiper.ID,
			Img_url:   reqswiper.Imgurl,
			Shop_name: reqswiper.ShopName,
		}
		swipers = append(swipers, swiper)
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "查询成功", "data": swipers})

}

func UpdateSwiper(c *gin.Context) {
	type reqswiper struct {
		ID       uint   `json:"id" bind:"required"`
		Imgurl   string `json:"img_url" bind:"required"`
		ShopName string `json:"shop_name" bind:"required"`
	}
	var request reqswiper
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error("参数不完整")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	fmt.Println(request)
	db := common.GetDB()
	if err := db.Model(&model.Swiper{}).Where("id=?", request.ID).Update(map[string]interface{}{"imgurl": request.Imgurl, "shop_name": request.ShopName}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "修改失败"})
	}

	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "修改成功"})
}

func DeleteSwiper(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		zap.L().Error("参数不全")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不全"})
	}
	db := common.GetDB()
	if err := db.Where("id=?", id).Delete(&model.Swiper{}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "删除失败"})
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "删除成功"})

}

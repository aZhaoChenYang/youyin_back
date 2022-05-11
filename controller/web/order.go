package web

import (
	"github.com/gin-gonic/gin"
	"time"
	"youyin/model"
	"youyin/response"
)

// AddOrder 添加订单
func AddOrder(c *gin.Context) {
	type OrderFrom struct {
		ScriptId int    `json:"script_id" binding:"required"`
		ShopId   int    `json:"shop_id" binding:"required"`
		Time     string `json:"time" binding:"required"`
	}
	var orderFrom OrderFrom
	if err := c.BindJSON(&orderFrom); err != nil {
		response.ParamError(c, err)
		return
	}
	Time, err := time.Parse("2006-01-02 15:04:05", orderFrom.Time)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	order := &model.Order{
		ScriptId: uint(orderFrom.ScriptId),
		ShopId:   uint(orderFrom.ShopId),
		DateTime: Time,
	}
	err = order.Create()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// 获取订单列表
func GetOrderList(c *gin.Context) {
	orderList, err := (&model.Order{}).GetList()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, orderList)
}

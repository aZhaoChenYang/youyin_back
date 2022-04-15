package app

import (
	"github.com/gin-gonic/gin"
	"youyin/model"
	"youyin/response"
)

func GetShopList(c *gin.Context) {
	shop, err := (&model.Shop{}).GetShop()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, shop)
}

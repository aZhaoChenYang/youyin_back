package web

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"youyin/model"
	"youyin/response"
)

func GetUserList(c *gin.Context) {
	list, err := (&model.User{}).GetList()
	if err != nil {
		return
	}
	response.Success(c, list)
}

func UpdateUser(c *gin.Context) {
	var quest struct {
		Id  int `json:"id" binding:"required"`
		Vip int `json:"vip" binding:"required"`
	}
	if err := c.ShouldBind(&quest); err != nil {
		response.ParamError(c, err)
		return
	}
	user := model.User{
		Model: gorm.Model{
			ID: uint(quest.Id),
		},
		Vip: quest.Vip,
	}
	if err := (&user).Update(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"youyin/model"
	"youyin/response"
)

// AddPeople 添加人数
func AddPeople(c *gin.Context) {
	var people model.People
	if err := c.BindJSON(&people); err != nil {
		response.ParamError(c, err)
		return
	}
	if err := people.Add(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// DeletePeople 删除人数
func DeletePeople(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.ParamError(c, err)
		return
	}
	if err := (&model.People{ID: uint(id)}).Delete(); err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// GetPeopleList 查询人数列表
func GetPeopleList(c *gin.Context) {
	list, err := (&model.People{}).GetList()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, list)
}

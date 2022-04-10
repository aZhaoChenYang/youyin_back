package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

// AddPeople 添加人数
func AddPeople(c *gin.Context) {
	var people model.People
	if err := c.BindJSON(&people); err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := people.Add(); err != nil {
		zap.L().Error("添加人数失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeletePeople 删除人数
func DeletePeople(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	var people model.People
	people.ID = uint(id)
	if err := people.Delete(); err != nil {
		zap.L().Error("删除人数失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// GetPeopleList 查询人数列表
func GetPeopleList(c *gin.Context) {
	var people model.People
	list, err := people.GetList()
	if err != nil {
		zap.L().Error("查询人数列表失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

//根据查询字符串type获取配置
func GetSetting(c *gin.Context) {
	typeName, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "type不能为空")
		return
	}
	var setting model.Setting
	settings, err := setting.GetList(typeName)
	if err != nil {
		return
	}
	response.Success(c, settings)

}

// 修改配置
func UpdateSetting(c *gin.Context) {
	var settings []model.Setting
	err := c.BindJSON(&settings)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	for _, setting := range settings {
		err := setting.Update()
		if err != nil {
			response.Error(c, http.StatusBadRequest, "修改失败")
			return
		}
	}
	response.Success(c, "修改成功")
}

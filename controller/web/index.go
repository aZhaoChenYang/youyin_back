package web

import (
	"github.com/gin-gonic/gin"
	"youyin/model"
	"youyin/response"
)

func GetIndex(c *gin.Context) {
	userCount, err := (&model.User{}).GetCount()
	if err != nil {
		response.DbError(c, err)
		return
	}
	scriptCount, err := (&model.Script{}).GetCount()
	if err != nil {
		response.DbError(c, err)
		return
	}
	type Info struct {
		Types string `json:"types"`
		Num   int64  `json:"num"`
		Mess  string `json:"mess"`
	}
	info := []Info{
		{
			Types: "用户数",
			Num:   userCount,
			Mess:  "用户数",
		},
		{
			Types: "脚本数",
			Num:   scriptCount,
			Mess:  "脚本数",
		},
	}
	response.Success(c, info)
}

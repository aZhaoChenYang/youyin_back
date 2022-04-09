package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"youyin/model"
	"youyin/response"
)

func AddScript(c *gin.Context) {
	var script model.Jsonscript
	if err := c.BindJSON(&script); err != nil {
		zap.L().Error("bind json error", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}

	peoples, err := (&model.People{}).GetFromIDS(script.Peoples)
	if err != nil {
		zap.L().Error("get people error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取人员数量失败")
		return
	}

	tags, err := (&model.Tag{}).GetFromIDS(script.Tags)
	if err != nil {
		zap.L().Error("get tag error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取标签信息失败")
		return
	}
	scriptModel := model.Script{
		Name:     script.Name,
		ImgUrl:   script.ImgUrl,
		Describe: script.Describe,
		Time:     script.Time,
		Boys:     script.Boys,
		Girls:    script.Girls,
		Price1:   script.Price1,
		Price2:   script.Price2,
		ShopID:   script.ShopID,
		Tags:     tags,
		Peoples:  peoples,
	}

	if err := scriptModel.Create(); err != nil {
		zap.L().Error("add script error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "添加剧本失败")
		return
	}
	response.Success(c, "添加剧本成功")
}

func GetScriptList(c *gin.Context) {
	list, err := (&model.Script{}).GetList()
	if err != nil {
		zap.L().Error("获取剧本列表失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取剧本列表失败")
		return
	}
	response.Success(c, list)
}

func GetScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	script, err := (&model.Script{}).Get(uint(id))
	if err != nil {
		zap.L().Error("获取剧本信息失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取剧本信息失败")
		return
	}
	response.Success(c, script)
}

func DeleteScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}
	script := model.Script{}
	script.ID = uint(id)
	if err := script.Delete(); err != nil {
		zap.L().Error("删除剧本失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "删除剧本失败")
		return
	}
	response.Success(c, "删除剧本成功")
}

func UpdateScript(c *gin.Context) {
	var script model.Jsonscript
	if err := c.BindJSON(&script); err != nil {
		zap.L().Error("bind json error", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "参数不完整")
		return
	}

	peoples, err := (&model.People{}).GetFromIDS(script.Peoples)
	if err != nil {
		zap.L().Error("get people error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取人员数量失败")
		return
	}

	tags, err := (&model.Tag{}).GetFromIDS(script.Tags)
	if err != nil {
		zap.L().Error("get tag error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "获取标签信息失败")
		return
	}
	scriptModel := model.Script{
		ID:       script.ID,
		Name:     script.Name,
		ImgUrl:   script.ImgUrl,
		Describe: script.Describe,
		Time:     script.Time,
		Boys:     script.Boys,
		Girls:    script.Girls,
		Price1:   script.Price1,
		Price2:   script.Price2,
		ShopID:   script.ShopID,
		Tags:     tags,
		Peoples:  peoples,
	}

	if err := scriptModel.Update(); err != nil {
		zap.L().Error("update script error", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "修改剧本失败")
		return
	}
	response.Success(c, "修改剧本成功")
}

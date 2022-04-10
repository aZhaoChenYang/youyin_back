package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"youyin/common"
	"youyin/model"
	"youyin/response"
)

func Login(c *gin.Context) {
	type LoginForm struct {
		Code string `form:"code" json:"code" binding:"required"`
	}
	var query LoginForm
	err := c.BindJSON(&query)
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusBadRequest, "参数不全")
		return
	}
	appid, err := (&model.Setting{}).GetValueByKey("wx_app_id")
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "获取appid失败")
		return
	}
	secret, err := (&model.Setting{}).GetValueByKey("wx_app_secret")
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "获取secret失败")
		return
	}
	requestUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, query.Code)
	fmt.Println(requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "获取session失败")
		return
	}
	defer resp.Body.Close()
	type result struct {
		Openid      string `json:"openid"`
		Unionid     string `json:"unionid"`
		Session_key string `json:"session_key"`
		ErrCode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
	}
	var res result
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "获取session失败")
		return
	}
	// 判断微信接口返回的是否是一个异常情况
	if res.ErrCode != 0 {
		zap.L().Error(res.Errmsg)
		response.Error(c, http.StatusInternalServerError, res.Errmsg)
		return
	}
	// 判断用户是否存在
	user := &model.User{
		OpenID: res.Openid,
	}
	if user.IsExist() {
		// 用户存在，更新session_key
		user.SessionKey = res.Session_key
		if err := user.Update(); err != nil {
			zap.L().Error(err.Error())
			response.Error(c, http.StatusInternalServerError, "更新session失败")
			return
		}
	} else {
		// 用户不存在，新增用户
		user.SessionKey = res.Session_key
		if err := user.Create(); err != nil {
			zap.L().Error(err.Error())
			response.Error(c, http.StatusInternalServerError, "新增用户失败")
			return
		}
	}

	// 生成token
	token, err := common.GenToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}
	response.Success(c, token)
}

func UpdateUserInfo(c *gin.Context) {
	type User struct {
		Avatar   string `form:"avatar" json:"avatar" binding:"required"`
		Nickname string `form:"nickname" json:"nickname" binding:"required"`
	}
	var query User
	err := c.BindJSON(&query)
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusBadRequest, "参数不全")
		return
	}
	userId, _ := c.Get("userId")
	id, err := strconv.Atoi(userId.(string))
	if err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "获取用户id失败")
		return
	}

	user := &model.User{
		Model:    gorm.Model{ID: uint(id)},
		Nickname: query.Nickname,
		Avatar:   query.Avatar,
	}
	fmt.Println(user)
	if err := user.Update(); err != nil {
		zap.L().Error(err.Error())
		response.Error(c, http.StatusInternalServerError, "更新用户信息失败")
		return
	}
	response.Success(c, "更新成功")
}

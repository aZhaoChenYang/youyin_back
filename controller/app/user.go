package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"strings"
	"youyin/common"
	"youyin/model"
	"youyin/response"
)

func Login(c *gin.Context) {

	type LoginForm struct {
		Code     string `form:"code" json:"code" binding:"required"`
		UserInfo struct {
			AvatarUrl string `json:"avatarUrl" binding:"required"`
			NickName  string `json:"nickName" binding:"required"`
		} `json:"userInfo" binding:"required"`
	}
	var query LoginForm
	err := c.BindJSON(&query)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	appid, err := (&model.Setting{}).GetValueByKey("wx_app_id")
	if err != nil {
		response.DbError(c, err)
		return
	}
	secret, err := (&model.Setting{}).GetValueByKey("wx_app_secret")
	if err != nil {
		response.DbError(c, err)
		return
	}
	requestUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, query.Code)
	fmt.Println(requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		response.ThirdPartyError(c, err)
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
		response.ThirdPartyError(c, err)
		return
	}
	// 判断微信接口返回的是否是一个异常情况
	if res.ErrCode != 0 {
		response.ThirdPartyError(c, fmt.Errorf("%s", res.Errmsg))
		return
	}
	// 判断用户是否存在
	user := &model.User{
		OpenID:     res.Openid,
		SessionKey: res.Session_key,
		Nickname:   query.UserInfo.NickName,
		Avatar:     query.UserInfo.AvatarUrl,
	}
	if user.IsExist() {
		// 用户存在，更新session_key
		if err := user.Update(); err != nil {
			response.DbError(c, err)
			return
		}
	} else {
		// 用户不存在，新增用户
		if err := user.Create(); err != nil {
			response.DbError(c, err)
			return
		}
	}
	vip, err := (&model.Setting{}).GetVipSetting(user.Vip)
	if err != nil {
		response.DbError(c, err)
		return
	}
	// 生成token
	token, err := common.GenToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		response.GenTokenError(c, err)
		return
	}
	response.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
			"phone":    user.Phone,
			"vip":      vip,
		},
	})
}

func UpdateUserPhone(c *gin.Context) {
	type UpdateUserPhoneForm struct {
		Code string `json:"code" binding:"required"`
	}
	var query UpdateUserPhoneForm
	err := c.BindJSON(&query)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	secret, err := (&model.Setting{}).GetValueByKey("wx_app_secret")
	if err != nil {
		response.DbError(c, err)
		return
	}
	appid, err := (&model.Setting{}).GetValueByKey("wx_app_id")
	if err != nil {
		response.DbError(c, err)
		return
	}
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, secret))
	if err != nil {
		response.ThirdPartyError(c, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.ThirdPartyError(c, err)
			return
		}
	}(resp.Body)
	type resultToken struct {
		Access_token string `json:"access_token"`
	}
	var resToken resultToken
	decoderToken := json.NewDecoder(resp.Body)
	if err := decoderToken.Decode(&resToken); err != nil {
		response.ThirdPartyError(c, err)
		return
	}
	resp, err = http.Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token="+resToken.Access_token, "application/json", strings.NewReader(`{"code":"`+query.Code+`"}`))
	if err != nil {
		response.ThirdPartyError(c, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.ThirdPartyError(c, err)
			return
		}
	}(resp.Body)
	type result struct {
		Errcode    int    `json:"errcode"`
		Errmsg     string `json:"errmsg"`
		Phone_info struct {
			PhoneNumber     string `json:"phoneNumber"`
			PurePhoneNumber string `json:"purePhoneNumber"`
			CountryCode     string `json:"countryCode"`
			Watermark       struct {
				Timestamp int64 `json:"timestamp"`
			} `json:"watermark"`
		} `json:"phone_info"`
	}
	var res result
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		response.ThirdPartyError(c, err)
		return
	}
	if res.Errcode != 0 {
		response.ThirdPartyError(c, fmt.Errorf("%s", res.Errmsg))
		return
	}
	userId, exist := c.Get("userId")
	if exist == false {
		response.ParamError(c, fmt.Errorf("qwerty"))
		return
	}
	id, err := strconv.Atoi(userId.(string))
	if err != nil {
		response.ParamError(c, err)
		return
	}

	err = (&model.User{
		Phone: res.Phone_info.PhoneNumber,
		Model: gorm.Model{ID: uint(id)},
	}).Update()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

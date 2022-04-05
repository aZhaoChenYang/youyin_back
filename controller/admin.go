package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)

func AddAdmin(c *gin.Context) {
	var request model.Admin
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error(err.Error())
		response.Response(c, http.StatusBadRequest, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数有误"})
		return
	}
	username := request.Username
	password := request.Password
	nickname := request.Nickname
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error(err.Error())
		response.Response(c, http.StatusUnprocessableEntity, gin.H{"errno": reqstatus.THIRDERR, "errmsg": "加密错误"})
		return
	}
	if nickname == "" {
		nickname = username
	}
	db := common.GetDB()
	admin := model.Admin{
		Username: username,
		Password: string(hashPassword),
		Nickname: nickname,
	}
	if err = db.Create(&admin).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "添加数据库失败"})
		return
	}

	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "添加成功"})
}

func Login(c *gin.Context) {
	var request model.Admin
	if err := c.BindJSON(&request); err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数有误"})
		return
	}
	username := request.Username
	password := request.Password

	db := common.GetDB()
	var admin model.Admin

	if err := db.Where("username = ?", username).First(&admin).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "数据库查询失败"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "用户名密码错误"})
		return
	}

	token, err := common.GenToken(admin.Username, int(admin.ID))
	if err != nil {
		response.Success(c, gin.H{"errno": reqstatus.THIRDERR, "errmsg": "生成token失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "登录成功", "data": gin.H{"user": admin.Username, "token": token}})

}

func GetAdmin(c *gin.Context) {
	db := common.GetDB()
	var admin []model.Admin
	if err := db.Select("id,username,nickname").Find(&admin).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "数据库查询失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "查询成功", "data": admin})
}

func DeleteAdmin(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		zap.L().Error("参数不足")
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数错误"})
		return
	}

	db := common.GetDB()
	if err := db.Where("username=?", username).Delete(&model.Admin{}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "数据库查询失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "删除成功"})
}

func UpdateAdmin(c *gin.Context) {
	var request model.Admin
	err := c.BindJSON(&request)
	if err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数有误"})
		return
	}
	username := request.Username
	password := request.Password
	nickname := request.Nickname
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.THIRDERR, "errmsg": "加密错误"})
		return
	}
	if nickname == "" {
		nickname = username
	}
	db := common.GetDB()
	if err := db.Model(&model.Admin{}).Where("username=?", username).Update(map[string]interface{}{"nickname": nickname, "password": hashPassword}).Error; err != nil {
		zap.L().Error(err.Error())
		response.Success(c, gin.H{"errno": reqstatus.DBERR, "errmsg": "修改失败"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "修改成功"})
}

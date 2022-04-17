package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"youyin/common"
	"youyin/model"
	"youyin/response"
)

// AddAdmin 添加管理员
func AddAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("加密密码失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	admin.Password = string(hashPassword)
	//判断nickname是否为空，是就将username赋值给nickname
	if admin.Nickname == "" {
		admin.Nickname = admin.Username
	}

	err = admin.Create()
	if err != nil {
		zap.L().Error("添加管理员失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// 获取全部管理员
func GetAllAdmin(c *gin.Context) {
	var admin model.Admin
	admins, err := admin.GetAll()
	if err != nil {
		zap.L().Error("获取管理员失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, admins)
}

// 更新管理员信息
func UpdateAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("加密密码失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	admin.Password = string(hashPassword)
	err = admin.Update()
	if err != nil {
		zap.L().Error("更新管理员失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// DeleteAdmin 根据查询参数ID删除管理员
func DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	admin := model.Admin{ID: uint(id)}
	err = admin.Delete()
	if err != nil {
		zap.L().Error("删除管理员失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

// 登录
func Login(c *gin.Context) {
	var admin model.Admin
	err := c.ShouldBind(&admin)
	if err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	err = admin.Login()
	if err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	//生成token
	token, err := common.GenToken(strconv.Itoa(int(admin.ID)))
	if err != nil {
		zap.L().Error("生成token失败", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}

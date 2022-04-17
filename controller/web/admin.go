package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"youyin/common"
	"youyin/model"
	"youyin/response"
)

// AddAdmin 添加管理员
func AddAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		response.EncryptionError(c, err)
		return
	}
	admin.Password = string(hashPassword)
	//判断nickname是否为空，是就将username赋值给nickname
	if admin.Nickname == "" {
		admin.Nickname = admin.Username
	}

	err = admin.Create()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// 获取全部管理员
func GetAllAdmin(c *gin.Context) {
	admins, err := (&model.Admin{}).GetAll()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, admins)
}

// 更新管理员信息
func UpdateAdmin(c *gin.Context) {
	var admin model.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		response.EncryptionError(c, err)
		return
	}
	admin.Password = string(hashPassword)
	err = admin.Update()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// DeleteAdmin 根据查询参数ID删除管理员
func DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.ParamError(c, err)
		return
	}
	err = (&model.Admin{ID: uint(id)}).Delete()
	if err != nil {
		response.DbError(c, err)
		return
	}
	response.Success(c, nil)
}

// 登录
func Login(c *gin.Context) {
	var admin model.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		response.ParamError(c, err)
		return
	}
	err = admin.Login()
	if err != nil {
		response.DbError(c, err)
		return
	}
	//生成token
	token, err := common.GenToken(strconv.Itoa(int(admin.ID)))
	if err != nil {
		response.GenTokenError(c, err)
		return
	}
	response.Success(c, gin.H{"token": token})
}

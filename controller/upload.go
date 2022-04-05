package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"go.uber.org/zap"
	"mime/multipart"
	"youyin/common"
	"youyin/common/reqstatus"
	"youyin/model"
	"youyin/response"
)
import (
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("img")
	if err != nil {
		response.Success(c, gin.H{"errno": reqstatus.PARAMERR, "errmsg": "参数不完整"})
		return
	}
	fileSize := fileHeader.Size
	url, code := uploadFile(file, fileSize)
	if code != 0 {
		response.Success(c, gin.H{"errno": reqstatus.THIRDERR, "errmsg": "第三方库调用错误"})
		return
	}
	response.Success(c, gin.H{"errno": reqstatus.OK, "errmsg": "上传成功", "data": url})

}

func uploadFile(file multipart.File, fileSize int64) (string, int) {
	accessKey, secretKey, ImgUrl, e := parseConfig()
	if e != 0 {
		return "", e
	}
	putPolicy := storage.PutPolicy{
		Scope: "youyinjushe",
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		zap.L().Error(err.Error())
		return "", reqstatus.THIRDERR
	}
	url := ImgUrl + ret.Key
	return url, reqstatus.OK
}

func parseConfig() (string, string, string, int) {
	var set []model.Setting
	db := common.GetDB()
	if err := db.Where("type=3").Find(&set).Error; err != nil {
		zap.L().Error(err.Error())
		return "", "", "", reqstatus.DBERR
	}
	var (
		accessKey, secretKey, url string
	)
	for _, setting := range set {
		if setting.Key == "ak" {
			accessKey = setting.Value
		} else if setting.Key == "sk" {
			secretKey = setting.Value
		} else if setting.Key == "url" {
			url = setting.Value
		}
	}
	return accessKey, secretKey, url, reqstatus.OK
}

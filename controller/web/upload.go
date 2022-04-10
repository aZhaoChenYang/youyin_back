package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"youyin/model"
	"youyin/response"
)
import (
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("img")
	if err != nil {
		zap.L().Error("参数不完整", zap.Error(err))
		response.Error(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	fileSize := fileHeader.Size
	url, code := uploadFile(file, fileSize)
	if code != 0 {
		zap.L().Error("上传失败", zap.Error(err))
		response.Error(c, http.StatusServiceUnavailable, "上传失败")
		return
	}
	response.Success(c, url)
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
		return "", http.StatusInternalServerError
	}
	url := ImgUrl + ret.Key
	return url, 0
}

func parseConfig() (string, string, string, int) {
	var set []model.Setting
	db := model.GetDB()
	if err := db.Where("type=3").Find(&set).Error; err != nil {
		zap.L().Error(err.Error())
		return "", "", "", http.StatusInternalServerError
	}
	var (
		accessKey, secretKey, url string
	)
	for _, setting := range set {
		if setting.AppKey == "ak" {
			accessKey = setting.Value
		} else if setting.AppKey == "sk" {
			secretKey = setting.Value
		} else if setting.AppKey == "url" {
			url = setting.Value
		}
	}
	return accessKey, secretKey, url, 0
}

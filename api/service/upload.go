package service

import (
	"fmt"
	"go-admin/common/config"
	"go-admin/common/result"
	"go-admin/util"
	"log"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IUploadService interface {
	Upload(c *gin.Context)
}

type UploadServiceImpl struct{}

func (u UploadServiceImpl) Upload(c *gin.Context) {
	log.Println("UploadDir =", config.Config.Imgupload.UploadDir)
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.UPLOADFILEFAILED), result.ApiCode.GetMessage(result.ApiCode.UPLOADFILEFAILED))
		return
	}
	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		config.Config.Imgupload.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	err = util.CreateDir(filePath)
	if err != nil {
		return
	}
	fullPath := filePath + "/" + fileName
	c.SaveUploadedFile(file, fullPath)
	result.Success(c, config.Config.Imgupload.Imghost+fullPath)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}

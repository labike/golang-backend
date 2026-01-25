package controller

import (
	"go-admin/api/service"

	"github.com/gin-gonic/gin"
)

// @Summary 上传
// @Produce json
// @Description 上传
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} result.Result
// @router /api/upload [post]
// @Security ApiKeyAuth
func Upload(c *gin.Context) {
	service.UploadService().Upload(c)
}

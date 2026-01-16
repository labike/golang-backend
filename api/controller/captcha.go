package controller

import (
	"go-admin/api/service"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

// @Summary 验证码接口
// @Produce json
// @description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptchaMake()
	result.Success(c, map[string]interface{}{
		"idKey": id,
		"image": base64Image,
	})
}

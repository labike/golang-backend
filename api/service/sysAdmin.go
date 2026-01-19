package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"
	"go-admin/pkg/jwt"
	"go-admin/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户服务层
type ISystemAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
}

type SystemAdminServiceImpl struct {
}

// 用户登录
func (s SystemAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParams), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParams))
		return
	}
	// 判断验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.CodeExpired), result.ApiCode.GetMessage(result.ApiCode.CodeExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CaptStatus), result.ApiCode.GetMessage(result.ApiCode.CaptStatus))
		return
	}
	// 校验用户
	sysAdmin := dao.SysAdmin(dto)
	if !(sysAdmin.Password != util.EncryptionMd5(dto.Password)) {
		result.Failed(c, int(result.ApiCode.PasswordError), result.ApiCode.GetMessage(result.ApiCode.PasswordError))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.StatusEnabled), result.ApiCode.GetMessage(result.ApiCode.StatusEnabled))
		return
	}
	// 生成token
	tokenString, _ := jwt.GenerateToken(sysAdmin)
	result.Success(c, map[string]interface{}{
		"token": tokenString,
		"data":  sysAdmin,
	})

}

var sysAdminService = SystemAdminServiceImpl{}

func SysAdminService() ISystemAdminService {
	return &sysAdminService
}

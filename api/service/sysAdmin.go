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
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetAdminInfoById(c *gin.Context, id int)
	UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	DeleteSysAdmin(c *gin.Context, dto entity.SysAdminIdDto)
	UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetAdminList(
		c *gin.Context,
		PageNum,
		PageSize int,
		Username,
		Status,
		BeginTime,
		EndTime string)
	UpdateUser(c *gin.Context, dto entity.UpdateUserDto)
	UpdateUserPassword(c *gin.Context, dto entity.UpdateUserPasswordDto)
}

type SystemAdminServiceImpl struct {
}

func (s SystemAdminServiceImpl) UpdateUserPassword(c *gin.Context, dto entity.UpdateUserPasswordDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.RESETPASSWORDFAILED), result.ApiCode.GetMessage(result.ApiCode.RESETPASSWORDFAILED))
		return
	}
	sysAdmin, _ := jwt.GetAdmin(c)
	dto.Id = sysAdmin.ID
	sysAdminExists := dao.GetSysAdminByUsername(sysAdmin.Username)
	if sysAdminExists.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.OLDPASSWORDERROR), result.ApiCode.GetMessage(result.ApiCode.OLDPASSWORDERROR))
		return
	}
	if dto.NewPassword != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.TWICEPASSWORDERROR), result.ApiCode.GetMessage(result.ApiCode.TWICEPASSWORDERROR))
		return
	}
	dto.NewPassword = util.EncryptionMd5(dto.NewPassword)
	sysAdminUpdatePwd := dao.UpdateUserPassword(dto)
	tokenString, _ := jwt.GenerateToken(sysAdminUpdatePwd)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdminUpdatePwd})
	return
}

func (s SystemAdminServiceImpl) UpdateUser(c *gin.Context, dto entity.UpdateUserDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.UPDATEUSERFAILED), result.ApiCode.GetMessage(result.ApiCode.UPDATEUSERFAILED))
		return
	}
	id, _ := jwt.GetAdminId(c)
	dto.Id = id
	result.Success(c, dao.UpdateUser(dto))
}

func (s SystemAdminServiceImpl) GetAdminList(c *gin.Context, PageNum, PageSize int, Username, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysAdmin, count := dao.GetAdminList(PageNum, PageSize, Username, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysAdmin,
	})
}

func (s SystemAdminServiceImpl) ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.ResetSysAdminPassword(dto)
	result.Success(c, true)
}

func (s SystemAdminServiceImpl) UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto) {
	dao.UpdateSysAdminStatus(dto)
	result.Success(c, true)
}

func (s SystemAdminServiceImpl) DeleteSysAdmin(c *gin.Context, dto entity.SysAdminIdDto) {
	dao.DeleteSysAdmin(dto)
	result.Success(c, true)
}

func (s SystemAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto) {
	result.Success(c, dao.UpdateSysAdmin(dto))
}

func (s SystemAdminServiceImpl) GetAdminInfoById(c *gin.Context, id int) {
	result.Success(c, dao.GetAdminInfoById(id))
}

func (s SystemAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	ok := dao.CreateSysAdmin(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.CREATEUSERFAILED), result.ApiCode.GetMessage(result.ApiCode.CREATEUSERFAILED))
		return
	}
	result.Success(c, true)
}

// 用户登录
func (s SystemAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParams), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParams))
		return
	}

	ip := c.ClientIP()

	// 判断验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		dao.CreateLoginlog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码已过期", 2)
		result.Failed(c, int(result.ApiCode.CodeExpired), result.ApiCode.GetMessage(result.ApiCode.CodeExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		dao.CreateLoginlog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码错误", 2)
		result.Failed(c, int(result.ApiCode.CaptStatus), result.ApiCode.GetMessage(result.ApiCode.CaptStatus))
		return
	}
	// 校验用户
	sysAdmin := dao.SysAdmin(dto)
	if !(sysAdmin.Password != util.EncryptionMd5(dto.Password)) {
		dao.CreateLoginlog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "账号或密码错误", 2)
		result.Failed(c, int(result.ApiCode.PasswordError), result.ApiCode.GetMessage(result.ApiCode.PasswordError))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		dao.CreateLoginlog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "账号已被禁用", 2)
		result.Failed(c, int(result.ApiCode.StatusEnabled), result.ApiCode.GetMessage(result.ApiCode.StatusEnabled))
		return
	}
	// 生成token
	tokenString, _ := jwt.GenerateToken(sysAdmin)
	dao.CreateLoginlog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "登陆成功", 1)

	//生成左侧菜单列表
	var leftMenuVo []entity.LeftMenuVo
	leftMenuList := dao.QueryLeftMenuList(sysAdmin.ID)
	for _, v := range leftMenuList {
		menuSvoList := dao.QueryMenuList(sysAdmin.ID, v.Id)
		item := entity.LeftMenuVo{}
		item.MenuVoList = menuSvoList
		item.Id = v.Id
		item.Icon = v.Icon
		item.MenuName = v.MenuName
		item.Url = v.Url
		leftMenuVo = append(leftMenuVo, item)
	}

	// 权限
	permissionsList := dao.QueryUserPermissionList(sysAdmin.ID)
	var strList = make([]string, 0)
	for _, v := range permissionsList {
		strList = append(strList, v.Value)
	}

	result.Success(c, map[string]interface{}{
		"token":        tokenString,
		"data":         sysAdmin,
		"leftMenuList": leftMenuVo,
		"permissions":  strList,
	})

}

var sysAdminService = SystemAdminServiceImpl{}

func SysAdminService() ISystemAdminService {
	return &sysAdminService
}

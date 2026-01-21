package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
}

type SysRoleServiceImpl struct{}

func (s SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto) {
	ok := dao.CreateSysRole(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.ADDROLEFAILED), result.ApiCode.GetMessage(result.ApiCode.ADDROLEFAILED))
		return
	}
	result.Success(c, true)
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}

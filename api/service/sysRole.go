package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
	GetSysRoleById(c *gin.Context, Id int)
	UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto)
	DeleteSysRole(c *gin.Context, dto entity.SysRoleIdDto)
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto)
	GetSysRoleList(
		c *gin.Context,
		PageNum,
		PageSize int,
		RoleName,
		Status,
		BeginTime,
		EndTime string)
	GetSysRoleVo(c *gin.Context)
	QueryRoleMenuIdList(c *gin.Context, Id int)
	AssignPermissions(c *gin.Context, menu entity.RoleMenu)
}

type SysRoleServiceImpl struct{}

func (s SysRoleServiceImpl) AssignPermissions(c *gin.Context, menu entity.RoleMenu) {
	result.Success(c, dao.AssignPermission(menu))
}

func (s SysRoleServiceImpl) QueryRoleMenuIdList(c *gin.Context, Id int) {
	roleMenuIdList := dao.QueryRoleMenuIdList(Id)
	var idList = make([]int, 0)
	for _, id := range roleMenuIdList {
		idList = append(idList, id.Id)
	}
	result.Success(c, idList)
}

func (s SysRoleServiceImpl) GetSysRoleVo(c *gin.Context) {
	result.Success(c, dao.GetSysRoleVo())
}

func (s SysRoleServiceImpl) GetSysRoleList(
	c *gin.Context,
	PageNum,
	PageSize int,
	RoleName,
	Status,
	BeginTime,
	EndTime string) {
	if PageNum <= 1 {
		PageNum = 1
	}
	if PageSize <= 1 {
		PageSize = 10
	}
	sysRole, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"list":     sysRole,
		"pageNum":  PageNum,
		"pageSize": PageSize,
	})
	return
}

func (s SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto) {
	ok := dao.UpdateSysRoleStatus(dto)
	if !ok {
		return
	}
	result.Success(c, true)
}

func (s SysRoleServiceImpl) DeleteSysRole(c *gin.Context, dto entity.SysRoleIdDto) {
	dao.DeleteSysRoleById(dto)
	result.Success(c, true)
}

func (s SysRoleServiceImpl) UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto) {
	sysRole := dao.UpdateSysRole(dto)
	result.Success(c, sysRole)
}

func (s SysRoleServiceImpl) GetSysRoleById(c *gin.Context, Id int) {
	sysRole := dao.GetSysRoleById(Id)
	result.Success(c, sysRole)
}

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

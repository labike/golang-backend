package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysMenuService interface {
	CreateSysMenu(c *gin.Context, menu entity.SysMenu)
	GetSysMenuVoList(c *gin.Context)
	GetSysMenuById(c *gin.Context, Id int)
	UpdateSysMenu(c *gin.Context, menu entity.SysMenu)
	DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto)
	GetSysMenuList(c *gin.Context, menuName string, menuStatus string)
}

type SysMenuServiceImpl struct{}

func (s SysMenuServiceImpl) GetSysMenuList(c *gin.Context, menuName string, menuStatus string) {
	result.Success(c, dao.GetSysMenuList(menuName, menuStatus))
}

func (s SysMenuServiceImpl) DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto) {
	ok := dao.DeleteSysMenu(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.DELETEMENUFAILED), result.ApiCode.GetMessage(result.ApiCode.DELETEMENUFAILED))
		return
	}
	result.Success(c, true)
}

func (s SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, menu entity.SysMenu) {
	result.Success(c, dao.UpdateSysMenu(menu))
}

func (s SysMenuServiceImpl) GetSysMenuById(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysMenuById(Id))
}

func (s SysMenuServiceImpl) GetSysMenuVoList(c *gin.Context) {
	result.Success(c, dao.GetSysMenuVoList())
}

func (s SysMenuServiceImpl) CreateSysMenu(c *gin.Context, sysMenu entity.SysMenu) {
	ok := dao.CreateSysMenu(sysMenu)
	if !ok {
		result.Failed(c, int(result.ApiCode.MENUEXISTS), result.ApiCode.GetMessage(result.ApiCode.MENUEXISTS))
		return
	}
	result.Success(c, true)
}

var sysMenuService = SysMenuServiceImpl{}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}

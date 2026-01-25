package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysMenu entity.SysMenu

// @Summary 创建菜单
// @Produce json
// @Desceription 创建菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/add [post]
// @Security ApiKeyAuth
func CreateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().CreateSysMenu(c, sysMenu)
}

// @Summary 查询菜单下拉列表
// @Produce json
// @Description 查询菜单下拉列表
// @Success 200 {object} result.Result
// @router /api/menu/vo/list [get]
// @Security ApiKeyAuth
func QuerySysMenuVoList(c *gin.Context) {
	service.SysMenuService().GetSysMenuVoList(c)
}

// @Summary 根据菜单id查询
// @Produce json
// @Description 根据菜单id查询
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/menu/info [get]
// @Security ApiKeyAuth
func QuerySysMenuById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysMenuService().GetSysMenuById(c, Id)
}

// @Summary 编辑菜单
// @Produce json
// @Description 编辑菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/update [put]
// @Security ApiKeyAuth
func UpdateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().UpdateSysMenu(c, sysMenu)
}

// @Summary 删除菜单
// @Produce json
// @Description 删除菜单
// @Param data body entity.SysMenuIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu/delete [delete]
// @Security ApiKeyAuth
func DeleteSysMenu(c *gin.Context) {
	dto := entity.SysMenuIdDto{}
	_ = c.BindJSON(&dto)
	service.SysMenuService().DeleteSysMenu(c, dto)
}

// @Summary 获取菜单列表
// @Produce json
// @Description 获取菜单列表
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} result.Result
// @router /api/menu/list [get]
// @Security ApiKeyAuth
func GetSysMenuList(c *gin.Context) {
	menuName := c.Query("menuName")
	menuStatus := c.Query("menuStatus")
	service.SysMenuService().GetSysMenuList(c, menuName, menuStatus)
}

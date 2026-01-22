package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysRole entity.SysRole

// @Summary 创建角色
// @Produce json
// @Description 创建角色
// @Param data body entity.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/add [post]
func CreateSysRole(c *gin.Context) {
	var dto entity.AddSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().CreateSysRole(c, dto)
}

// @Summary 根据角色Id查询
// @Produce json
// @Description 根据角色Id查询
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/role/info [get]
func GetSysRoleById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, Id)
}

// @Summary 编辑角色
// @Produce json
// @Description 编辑角色
// @Param data body entity.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/update [put]
func UpdateSysRole(c *gin.Context) {
	var dto entity.UpdateSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRole(c, dto)
}

// @Summary 删除角色
// @Produce json
// @Description 删除角色
// @Param data body entity.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/delete [delete]
func DeleteSysRole(c *gin.Context) {
	var dto entity.SysRoleIdDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().DeleteSysRole(c, dto)
}

// @Summary 编辑角色状态
// @Produce json
// @Description 编辑角色状态
// @Param data body entity.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/updateStatus [put]
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// @Summary 获取角色列表(分页)
// @Produce json
// @Description 获取角色列表(分页)
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param roleName query string false "角色名称"
// @Param status query string false "状态 ->启用 2->禁用"
// @Param beginTime query string false "创建时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/role/list [get]
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pagenNum"))
	PageSize, _ := strconv.Atoi(c.Query("pagenSize"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysRoleService().GetSysRoleList(
		c,
		PageNum,
		PageSize,
		RoleName,
		Status,
		BeginTime,
		EndTime)
}

// @Summary 角色下拉列表
// @Produce json
// @Description 角色下拉列表
// @Success 200 {object} result.Result
// @router /api/role/vo/list [get]
func GetSysRoleVo(c *gin.Context) {
	service.SysRoleService().GetSysRoleVo(c)
}

// @Summary 根据角色id查询菜单权限
// @Produce json
// @Description 根据角色id查询菜单权限
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/role/vo/idList [get]
func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().QueryRoleMenuIdList(c, Id)
}

// @Summary 分配权限
// @Produce json
// @Description 分配权限
// @Param data body entity.RoleMenu true "data"
// @Success 200 {object} result.Result
// @router /api/role/assignPermission [put]
func AssignPermission(c *gin.Context) {
	var RoleMenu entity.RoleMenu
	_ = c.BindJSON(&RoleMenu)
	service.SysRoleService().AssignPermissions(c, RoleMenu)
}

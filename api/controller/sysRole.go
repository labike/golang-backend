package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"

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

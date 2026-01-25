package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysDept entity.SysDept

// @Summary 查询部门列表
// @Produce json
// @Description 查询部门列表
// @Param deptName query string false "部门名称"
// @Param deptStatus query string false "部门状态"
// @Success 200 {object} result.Result
// @router /api/dept/list [get]
// @Security ApiKeyAuth
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("deptName")
	DeptStatus := c.Query("deptStatus")
	service.SysDeptService().GetSysDeptList(c, DeptName, DeptStatus)
}

// @Summary 新增部门
// @Produce json
// @Description 新增部门
// @Param data body entity.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/add [post]
// @Security ApiKeyAuth
func CreateSysDept(c *gin.Context) {
	_ = c.Bind(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// @Summary 根据部门id查询
// @Produce json
// @Description 根据部门id查询
// @Param id query int true "ID"
// @Success 200 {object} result.Result
// @router /api/dept/info [get]
// @Security ApiKeyAuth
func GetSysDeptById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysDeptService().GetSysDeptById(c, Id)
}

// @Summary 编辑部门
// @Produce json
// @Description 编辑部门
// @Param data body entity.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/update [put]
// @Security ApiKeyAuth
func UpdateSysDept(c *gin.Context) {
	var dpt entity.SysDept
	_ = c.Bind(&dpt)
	service.SysDeptService().UpdateSysDept(c, dpt)
}

// @Summary 根据部门id删除
// @Produce json
// @Description 根据部门id删除
// @Param data body entity.SysDeptDto true "data"
// @Success 200 {object} result.Result
// @router /api/dept/delete [delete]
// @Security ApiKeyAuth
func DeleteSysDeptById(c *gin.Context) {
	var dto entity.SysDeptDto
	_ = c.Bind(&dto)
	service.SysDeptService().DeleteSysDeptById(c, dto)
}

// @Summary 获取部门下拉列表
// @Produce json
// Description 部门下拉列表
// @Success 200 {object} result.Result
// @router /api/dept/vo/list [get]
// @Security ApiKeyAuth
func QueryDeptList(c *gin.Context) {
	service.SysDeptService().GetDeptList(c)
}

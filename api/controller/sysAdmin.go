package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.Bind(&dto)
	service.SysAdminService().Login(c, dto)
}

// @Summary 创建用户
// @Produce json
// @Description 创建用户
// @Param data body entity.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/add [post]
// @Security ApiKeyAuth
func CreateSysAdmin(c *gin.Context) {
	var dto entity.AddSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().CreateSysAdmin(c, dto)
}

// @Summary 根据id查找用户
// @Produce json
// @Description 根据id查找用户
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/admin/info [get]
// @Security ApiKeyAuth
func GetAdminInfo(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysAdminService().GetAdminInfoById(c, Id)
}

// @Summary 编辑用户
// @Produce json
// @Description 编辑用户
// @Param data body entity.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/update [put]
// @Security ApiKeyAuth
func UpdateSysAdmin(c *gin.Context) {
	var dto entity.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdmin(c, dto)
}

// @Summary 删除用户
// @Produce json
// @Description 删除用户
// @Param data body entity.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/delete [delete]
// @Security ApiKeyAuth
func DeleteSysAdmin(c *gin.Context) {
	var dto entity.SysAdminIdDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().DeleteSysAdmin(c, dto)
}

// @Summary 编辑用户状态
// @Produce json
// @Description 编辑用户状态
// @Param data body entity.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updateStatus [put]
// @Security ApiKeyAuth
func UpdateSysAdminStatus(c *gin.Context) {
	var dto entity.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminStatus(c, dto)
}

// @Summary 重置密码
// @Produce json
// @Description 重置密码
// @Param data body entity.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/resetPassword [put]
// @Security ApiKeyAuth
func UpdateSysAdminPassword(c *gin.Context) {
	var dto entity.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().ResetSysAdminPassword(c, dto)
}

// @Summary 分页获取用户列表
// @Produce json
// @Description 分页获取用户列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户名"
// @Param status query string false "状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/admin/list [get]
// @Security ApiKeyAuth
func GetAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysAdminService().GetAdminList(c, PageNum, PageSize, Username, Status, BeginTime, EndTime)
}

// @Summary 更新用户信息
// @Produce json
// @Description 更新用户信息
// @Param data body entity.UpdateUserDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonal [put]
// @Security ApiKeyAuth
func UpdateUser(c *gin.Context) {
	var dto entity.UpdateUserDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateUser(c, dto)
}

// @Summary 更新用户密码
// @Produce json
// @Description 更新用户密码
// @Param data body entity.UpdateUserPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonalPassword [put]
// @Security ApiKeyAuth
func UpdateUserPwd(c *gin.Context) {
	var dto entity.UpdateUserPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateUserPassword(c, dto)
}

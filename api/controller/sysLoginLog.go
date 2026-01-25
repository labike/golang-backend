package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 分页查询登陆记录
// @Produce json
// @Description 分页查询登陆记录
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户账号"
// @Param loginStatus query string false "登陆状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/list [get]
// @Security ApiKeyAuth
func GetSysLoginLogList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	LoginStatus := c.Query("loginStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysLoginLogService().GetSysLoginLogList(c, Username, LoginStatus, BeginTime, EndTime, PageNum, PageSize)
}

// @Summary 批量删除登录日志
// @Produce json
// @Description 批量删除登录日志
// @Param data body entity.DeleteLoginInfoDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/batch/delete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysLoginLog(c *gin.Context) {
	var dto entity.DeleteLoginInfoDto
	_ = c.Bind(&dto)
	service.SysLoginLogService().BatchDeleteSysLoginLog(c, dto)
}

// @Summary 删除某个登录日志
// @Produce json
// @Description 删除某个登录日志
// @Param data body entity.SysLoginInfoIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/delete [delete]
// @Security ApiKeyAuth
func DeleteSysLoginLog(c *gin.Context) {
	var dto entity.SysLoginInfoIdDto
	_ = c.Bind(&dto)
	service.SysLoginLogService().DeleteSysLoginLog(c, dto)
}

// @Summary 清空登录日志
// @Produce json
// @Description 清空登录日志
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/clean [delete]
// @Security ApiKeyAuth
func CleanLoginLog(c *gin.Context) {
	service.SysLoginLogService().ClearSysLoginLog(c)
}

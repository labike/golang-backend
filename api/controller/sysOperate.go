package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 分页查询操作记录
// @Produce json
// @Description 分页查询操作记录
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户账号"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/list [get]
// @Security ApiKeyAuth
func GetSysOperateList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysOperateService().GetSysOperateList(c, Username, BeginTime, EndTime, PageSize, PageNum)
}

// @Summary 批量删除操作日志
// @Produce json
// @Description 批量删除操作日志
// @Param data body entity.BatchSysOperateDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/batch/delete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysOperateLog(c *gin.Context) {
	var dto entity.BatchSysOperateDto
	_ = c.Bind(&dto)
	service.SysOperateService().BatchDeleteOperateLog(c, dto)
}

// @Summary 删除某个操作日志
// @Produce json
// @Description 删除某个操作日志
// @Param data body entity.SysOperateIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/delete [delete]
// @Security ApiKeyAuth
func DeleteSysOperateLog(c *gin.Context) {
	var dto entity.SysOperateIdDto
	_ = c.Bind(&dto)
	service.SysOperateService().DeleteOperateLog(c, dto)
}

// @Summary 清空操作日志
// @Produce json
// @Description 清空操作日志
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/clean [delete]
// @Security ApiKeyAuth
func CleanSysOperateLog(c *gin.Context) {
	service.SysOperateService().CleanSysOperateLog(c)
}

package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysLoginLogService interface {
	GetSysLoginLogList(
		c *gin.Context,
		Username,
		LoginStatus,
		BeginTime,
		EndTime string,
		PageNum,
		PageSize int)
	BatchDeleteSysLoginLog(c *gin.Context, dto entity.DeleteLoginInfoDto)
	DeleteSysLoginLog(c *gin.Context, dto entity.SysLoginInfoIdDto)
	ClearSysLoginLog(c *gin.Context)
}

type SysLoginLogServiceImpl struct{}

func (s SysLoginLogServiceImpl) ClearSysLoginLog(c *gin.Context) {
	dao.ClearSysLoginLog()
	result.Success(c, true)
}

func (s SysLoginLogServiceImpl) BatchDeleteSysLoginLog(c *gin.Context, dto entity.DeleteLoginInfoDto) {
	dao.BatchDeleteSysLoginLog(dto)
	result.Success(c, true)
}

func (s SysLoginLogServiceImpl) DeleteSysLoginLog(c *gin.Context, dto entity.SysLoginInfoIdDto) {
	dao.DeleteSysLoginLog(dto)
	result.Success(c, true)
}

func (s SysLoginLogServiceImpl) GetSysLoginLogList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageNum, PageSize int) {
	if PageNum < 1 {
		PageNum = 1
	}
	if PageSize < 1 {
		PageSize = 10
	}
	sysLoginInfo, count := dao.GetSysLoginLogList(Username, LoginStatus, BeginTime, EndTime, PageNum, PageSize)
	result.Success(c, map[string]interface{}{
		"count":    count,
		"list":     sysLoginInfo,
		"pageNum":  PageNum,
		"pageSize": PageSize,
	})
}

var sysLoginLogService = SysLoginLogServiceImpl{}

func SysLoginLogService() ISysLoginLogService {
	return &sysLoginLogService
}

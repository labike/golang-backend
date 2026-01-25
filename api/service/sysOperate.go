package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysOperateService interface {
	GetSysOperateList(
		c *gin.Context,
		Username,
		BeginTime,
		EndTime string,
		PageSize,
		PageNum int)
	DeleteOperateLog(c *gin.Context, dto entity.SysOperateIdDto)
	BatchDeleteOperateLog(c *gin.Context, dto entity.BatchSysOperateDto)
	CleanSysOperateLog(c *gin.Context)
}

type SysOperateServiceImpl struct{}

func (s SysOperateServiceImpl) CleanSysOperateLog(c *gin.Context) {
	dao.CleanSysOperateLog()
	result.Success(c, true)
}

func (s SysOperateServiceImpl) GetSysOperateList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageNum < 1 {
		PageNum = 1
	}
	if PageSize < 1 {
		PageSize = 10
	}
	sysOperate, count := dao.GetSysOperateList(Username, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{
		"list":     sysOperate,
		"count":    count,
		"PageNum":  PageNum,
		"PageSize": PageSize,
	})
}

func (s SysOperateServiceImpl) DeleteOperateLog(c *gin.Context, dto entity.SysOperateIdDto) {
	dao.DeleteOperateLog(dto)
	result.Success(c, true)
}

func (s SysOperateServiceImpl) BatchDeleteOperateLog(c *gin.Context, dto entity.BatchSysOperateDto) {
	dao.BatchDeleteOperateLog(dto)
	result.Success(c, true)
}

func (s SysOperateServiceImpl) CreateSysOperateLog(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

var sysOperateService = SysOperateServiceImpl{}

func SysOperateService() ISysOperateService {
	return &sysOperateService
}

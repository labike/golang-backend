package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName, DeptStatus string)
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
	GetSysDeptById(c *gin.Context, Id int)
	UpdateSysDept(c *gin.Context, sysDept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptDto)
	GetDeptList(c *gin.Context)
}

type SysDeptServiceImpl struct{}

func (s SysDeptServiceImpl) GetDeptList(c *gin.Context) {
	result.Success(c, dao.GetDeptList())
}

func (s SysDeptServiceImpl) UpdateSysDept(c *gin.Context, sysDept entity.SysDept) {
	dept := dao.UpdateDysDept(sysDept)
	result.Success(c, dept)
}

func (s SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptDto) {
	ok := dao.DeleteSysDeptById(dto)
	if !ok {
		result.Failed(c, int(result.ApiCode.DEPTDELETEFAILED), result.ApiCode.GetMessage(result.ApiCode.DEPTDELETEFAILED))
		return
	}
	result.Success(c, true)
}

func (s SysDeptServiceImpl) GetSysDeptById(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysDepyById(Id))
}

func (s SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysdept entity.SysDept) {
	ok := dao.CreateSysDept(sysdept)
	if !ok {
		result.Failed(c, int(result.ApiCode.DEPTEXISTS), result.ApiCode.GetMessage(result.ApiCode.DEPTEXISTS))
		return
	}
	result.Success(c, true)
}

func (s SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName, DeptStatus string) {
	result.Success(c, dao.GetSysDeptList(DeptName, DeptStatus))
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}

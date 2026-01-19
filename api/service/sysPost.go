package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

// 服务层(调用dao层返回结果)
type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	GetPostById(c *gin.Context, Id int)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPostByIds(c *gin.Context, dto entity.DelSysPostDto)
}

type SysPostServiceImpl struct{}

func (s SysPostServiceImpl) BatchDeleteSysPostByIds(c *gin.Context, dto entity.DelSysPostDto) {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}

func (s SysPostServiceImpl) DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto) {
	dao.DeleteSysPostById(dto)
	result.Success(c, true)
}

func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost) {
	result.Success(c, dao.UpdateSysPost(sysPost))
}

func (s SysPostServiceImpl) GetPostById(c *gin.Context, Id int) {
	result.Success(c, dao.GetPostById(Id))
}

func (s SysPostServiceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysPost, total := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    total,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysPost,
	})
}

func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	ok := dao.CreateSysPost(sysPost)
	if !ok {
		result.Failed(c, int(result.ApiCode.POSTEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTEXISTS))
		return
	}
	result.Success(c, true)
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}

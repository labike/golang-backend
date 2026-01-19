package controller

// 调用service

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysPost entity.SysPost

// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// @Summary 分页查询岗位列表
// @Produce json
// @Description 分页查询岗位列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param postName query string false "岗位名称"
// @Param pageStatus query string false "岗位状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
func GetSysPostList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
}

// @Summary 根据岗位id查询
// @Produce json
// @Description 根据岗位id查询
// @Param id query int true "ID"
// @Success 200 {object} result.Result
// @router /api/post/info [get]
func GetPostById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysPostService().GetPostById(c, Id)
}

// @Summary 编辑岗位
// @Produce json
// @Description 编辑岗位
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/update [put]
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}

// @Summary 根据岗位id删除
// @Produce json
// @Description 根据岗位id删除
// @Param data body entity.SysPostIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/delete [delete]
func DeleteSysPostById(c *gin.Context) {
	// @Security ApiKeyAuth
	var dto entity.SysPostIdDto
	_ = c.BindJSON(&dto)
	service.SysPostService().DeleteSysPostById(c, dto)
}

// @Summary 批量删除岗位
// @Produce json
// @Description 批量删除岗位
// @Param data body entity.DelSysPostDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/batch/delete [delete]
func BatchDeleteSysPost(c *gin.Context) {
	var dto entity.DelSysPostDto
	_ = c.BindJSON(&dto)
	service.SysPostService().BatchDeleteSysPostByIds(c, dto)
}

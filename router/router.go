package router

import (
	"go-admin/api/controller"
	"go-admin/common/config"
	"go-admin/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())

	register(router)

	urlPath := "/upload"
	diskPath := config.Config.Imgupload.UploadDir
	router.Static(urlPath, diskPath)
	//router.StaticFS(config.Config.Imgupload.UploadDir, http.Dir(config.Config.Imgupload.UploadDir))

	return router
}

func register(router *gin.Engine) {
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)

	// jwt鉴权
	jwt := router.Group("/api", middleware.AuthMiddleware(), middleware.LogMiddleware())
	{
		jwt.POST("/post/add", controller.CreateSysPost)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/info", controller.GetPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)
		jwt.GET("/dept/list", controller.GetSysDeptList)
		jwt.POST("/dept/add", controller.CreateSysDept)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/vo/list", controller.QueryDeptList)
		jwt.POST("/menu/add", controller.CreateSysMenu)
		jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.GET("/menu/info", controller.QuerySysMenuById)
		jwt.PUT("/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/menu/delete", controller.DeleteSysMenu)
		jwt.GET("/menu/list", controller.GetSysMenuList)
		jwt.POST("/role/add", controller.CreateSysRole)
		jwt.GET("/role/info", controller.GetSysRoleById)
		jwt.PUT("/role/update", controller.UpdateSysRole)
		jwt.DELETE("/role/delete", controller.DeleteSysRole)
		jwt.PUT("/role/updateStatus", controller.UpdateSysRoleStatus)
		jwt.GET("/role/list", controller.GetSysRoleList)
		jwt.GET("/role/vo/list", controller.GetSysRoleVo)
		jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/role/assignPermission", controller.AssignPermission)
		jwt.POST("/admin/add", controller.CreateSysAdmin)
		jwt.GET("/admin/info", controller.GetAdminInfo)
		jwt.PUT("/admin/update", controller.UpdateSysAdmin)
		jwt.DELETE("/admin/delete", controller.DeleteSysAdmin)
		jwt.PUT("/admin/updateStatus", controller.UpdateSysAdminStatus)
		jwt.PUT("/admin/resetPassword", controller.UpdateSysAdminPassword)
		jwt.GET("/admin/list", controller.GetAdminList)
		jwt.POST("/upload", controller.Upload)
		jwt.PUT("/admin/updatePersonal", controller.UpdateUser)
		jwt.PUT("/admin/updatePersonalPassword", controller.UpdateUserPwd)
		jwt.GET("/sysLoginInfo/list", controller.GetSysLoginLogList)
		jwt.DELETE("/sysLoginInfo/delete", controller.DeleteSysLoginLog)
		jwt.DELETE("/sysLoginInfo/batch/delete", controller.BatchDeleteSysLoginLog)
		jwt.DELETE("/sysLoginInfo/clean", controller.CleanLoginLog)
		jwt.GET("/sysOperationLog/list", controller.GetSysOperateList)
		jwt.DELETE("/sysOperationLog/delete", controller.DeleteSysOperateLog)
		jwt.DELETE("/sysOperationLog/batch/delete", controller.BatchDeleteSysOperateLog)
		jwt.DELETE("/sysOperationLog/clean", controller.CleanSysOperateLog)
	}
}

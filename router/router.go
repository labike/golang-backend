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
	router.POST("/api/post/add", controller.CreateSysPost)
	router.GET("/api/post/list", controller.GetSysPostList)
	router.GET("/api/post/info", controller.GetPostById)
	router.PUT("/api/post/update", controller.UpdateSysPost)
	router.DELETE("/api/post/delete", controller.DeleteSysPostById)
	router.DELETE("/api/post/batch/delete", controller.BatchDeleteSysPost)
	router.PUT("/api/post/updateStatus", controller.UpdateSysPostStatus)
	router.GET("/api/post/vo/list", controller.QuerySysPostVoList)
	router.GET("/api/dept/list", controller.GetSysDeptList)
	router.POST("/api/dept/add", controller.CreateSysDept)
	router.GET("/api/dept/info", controller.GetSysDeptById)
	router.PUT("/api/dept/update", controller.UpdateSysDept)
	router.DELETE("/api/dept/delete", controller.DeleteSysDeptById)
	router.GET("/api/dept/vo/list", controller.QueryDeptList)
	router.POST("/api/menu/add", controller.CreateSysMenu)
	router.GET("/api/menu/vo/list", controller.QuerySysMenuVoList)
	router.GET("/api/menu/info", controller.QuerySysMenuById)
	router.PUT("/api/menu/update", controller.UpdateSysMenu)
	router.DELETE("/api/menu/delete", controller.DeleteSysMenu)
	router.GET("/api/menu/list", controller.GetSysMenuList)
	router.POST("/api/role/add", controller.CreateSysRole)
}

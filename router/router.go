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
}

package router

import (
	"go-admin/common/config"
	"go-admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.Imgupload.UploadDir, http.Dir(config.Config.Imgupload.UploadDir))
	router.Use(middleware.Logger())
	register(router)
	return router
}

func register(router *gin.Engine) {

}

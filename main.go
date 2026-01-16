package main

import (
	"context"
	"go-admin/common/config"
	"go-admin/pkg/db"
	log2 "go-admin/pkg/log"
	"go-admin/pkg/redis"
	router2 "go-admin/router"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	log := log2.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router2.InitRouter()
	ser := &http.Server{
		Addr:    ":" + config.Config.Server.Address,
		Handler: router,
	}

	go func() {
		if err := ser.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s\n", err)
		}
		log.Info("%s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ser.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

func init() {
	db.SetupDB()
	redis.SetupRedis()
}

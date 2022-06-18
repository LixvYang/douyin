// Package main provides the main function
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tiktok/model"
	"tiktok/pkg/middleware"
	"tiktok/pkg/redis"
	"tiktok/routes"
	"tiktok/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	signalch := make(chan os.Signal, 1)
	
	// 启动MySQL数据库
	model.InitDb()
	// 启动Redis数据库
	redis.InitRedis()
	// Gin框架初始化
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Println(err)
	}
	gin.SetMode(utils.AppMode)
  go routes.InitRouter(
		// cores
		r,
		middleware.Cors(),
		middleware.Log(),
		gin.Recovery(),
		gin.Logger(),
	)

	go func() {
		if err := r.Run(utils.HttpPort); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	/** 
	 * 监听信号
	*/
	signal.Notify(signalch, os.Interrupt, syscall.SIGTERM, os.Kill)
	signalType := <-signalch
	signal.Stop(signalch)
	log.Printf("Os Signal <%s>", signalType)
	log.Printf("Exit command received. Exiting...")
}

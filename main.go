// Package main provides the main function
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tiktok/model"
	"tiktok/routes"
	"tiktok/utils"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	signalch := make(chan os.Signal, 1)
	
	model.InitDb()
	redis.InitRedis()
	r := gin.New()
	gin.SetMode(utils.AppMode)
  go routes.InitRouter(
		// cores
		r,
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

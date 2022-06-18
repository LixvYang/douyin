// Package routes provides the routes for the application.
package routes

import (
	"net/http"
	"tiktok/api/v1/comment"
	"tiktok/api/v1/favorite"
	"tiktok/api/v1/follow"
	"tiktok/api/v1/user"
	"tiktok/api/v1/video"

	"github.com/gin-gonic/gin"
	"tiktok/pkg/middleware"
)

func InitRouter(r *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	r.MaxMultipartMemory = 32 << 21 // 1024M
	/**
	  Use Middlewares
		中间件
	*/
	gin.ForceConsoleColor()
	r.Use(gin.Recovery(), gin.Logger(), middleware.Cors(), middleware.Log())

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	r.Static("/public", "./public")
	rDouyin := r.Group("douyin")
	rDouyin.GET("/feed", video.GetVideoList)
	// 视频
	{
		rDouyin.POST("/publish/action", video.AddVideo)
		rDouyin.GET("/publish/list", video.GetUserPublishList)
	}
	// 用户
	{
		rDouyin.GET("/user", user.GetUserById)
		rDouyin.POST("/user/register", user.AddUser)
		rDouyin.POST("/user/login", user.Login)
	}
	// 点赞操作
	{
		rDouyin.POST("/favorite/action", favorite.AddFavorite)
		rDouyin.GET("/favorite/list", favorite.GetFavoriteList)
	}
	// 评论操作
	{
		rDouyin.POST("/comment/action", comment.AddComment)
		rDouyin.GET("/comment/list", comment.GetCommentList)
	}
	// 关注操作
	{
		rDouyin.POST("/relation/action/", follow.AddFollow)
		rDouyin.GET("/relation/follow/list/", follow.GetFollowList)
		rDouyin.GET("/relation/follower/list/", follow.GetFollowerList)
	}
	return r
}

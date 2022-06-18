// Package video provides video for api.
package video

import (
	"fmt"
	"net/http"
	"path/filepath"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
)

// 上传视频
// FIXME 尚未解决Cover_url 问题
// 尚未解决
func AddVideo(c *gin.Context) {
	// 未获取到token
	token, exist := c.GetPostForm("token")
	if !exist {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrUserLogin.StatusCode,
			StatusMsg:  errmsg.ErrUserLogin.StatusMsg,
		})
		return
	}

	// 根据token未找到用户
	rcGet := redis.RCGet(token)
	if rcGet == nil {
		fmt.Println("未找到用户")
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrUserLogin.StatusCode,
			StatusMsg:  errmsg.ErrUserLogin.StatusMsg,
		})
		return
	}

	// 视频数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrGetFile.StatusCode,
			StatusMsg:  errmsg.ErrGetFile.StatusMsg,
		})
		return
	}

	filename := filepath.Base(data.Filename)
	// fmt.Println(filename[:len(filename)-len(filepath.Ext(filename))]) // 文件名
	// fmt.Println(filepath.Ext(filename)) // 文件后缀
	// 只允许上传 .mp4 类型文件
	if filepath.Ext(filename) != ".mp4" {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrFileType.StatusCode,
			StatusMsg:  errmsg.ErrFileType.StatusMsg,
		})
		return
	}
	userId, _ := rcGet.Int64()
	// 创建视频 public 文件夹
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrSaveFile.StatusCode,
			StatusMsg:  errmsg.ErrSaveFile.StatusMsg,
		})
		return
	}
	title, _ := c.GetPostForm("token")
	updateVideo := &model.Video{
		Author:     userId,
		PlayUrl:    "http://121.196.105.17:8080" + "/golang/douyin/public/" + finalName,
		CoverUrl:   "",
		CreateTime: time.Now().Unix(),
		Title:      title,
	}
	if code := model.AddVideo(updateVideo); code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ErrDatabase.StatusCode,
			StatusMsg:  errmsg.ErrDatabase.StatusMsg,
		})
		return
	}
	c.JSON(http.StatusOK, common.Response{
		StatusCode: errmsg.OK.StatusCode,
		StatusMsg:  errmsg.OK.StatusMsg,
	})

}

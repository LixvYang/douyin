// Package favorite provides favorite for api.
package favorite

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)

// 通过user_id/video_id/action/type 点赞
func AddFavorite(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	// 检查用户是否存在
	userId, _ := redis.RCGet(token).Int64()
	fmt.Println(userId)
	if userId == 0 {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.ERROR.StatusCode,
			StatusMsg:  errmsg.ERROR.StatusMsg,
		})
		return
	}
	var data model.UsersLikeVideo
	data.UserId = userId
	data.VideoId = int64(video_id)
	// 点赞视频
	if action_type == 1 {
		data.IsLike = int64(action_type)
		code := model.AddFavorite(&data)
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "点赞失败")
		}

		code2 := model.SetFavoriteCount(int64(video_id), int64(action_type))
		if code2 != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "点赞总数修改失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		})
	} else if action_type == 2 {
		// 取消点赞
		code := model.DelFavorite(userId, int64(video_id))
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "取消点赞失败")
		}

		code2 := model.SetFavoriteCount(int64(video_id), int64(action_type))
		if code2 != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "点赞总数修改失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		})
	}
}

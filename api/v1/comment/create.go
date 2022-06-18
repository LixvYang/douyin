// Package comment provides comment for api.
package comment

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context)  {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	userid, _ := redis.RCGet(token).Int64()
	var comment model.Comment
	if action_type == 1 {
		comment_text := c.Query("comment_text")
		comment.FromUserId = int64(userid)
		comment.VideoId = int64(video_id)
		comment.Comment = comment_text
		comment.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		code := model.AddComment(&comment)
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "评论失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		})
	}	else if action_type == 2 {
		comment_id, _ := strconv.Atoi(c.Query("comment_id"))
		code := model.DelComment(int64(comment_id))
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "删除评论失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		})
	}
}
// Package comment provides comment for api.
package comment

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)
type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

func GetCommentList(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))

	userid, _ := redis.RCGet(token).Int64()

	modelCommentList, code := model.GetCommentListByVideoId(int64(video_id))
	if code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, "获取评论列表失败")
	}

	var commonCommentList []common.Comment
	for _, modelComment := range modelCommentList {
		var commonComment common.Comment
		commonComment.Id = modelComment.Id
		commonComment.Content = modelComment.Comment
		commonComment.CreateDate = modelComment.CreateDate

		commonCommentUser, _ := model.GetUserById(modelComment.FromUserId)
		commonComment.User.Id = commonCommentUser.Id
		commonComment.User.Name = commonCommentUser.Username
		commonComment.User.FollowCount = commonCommentUser.FollowCount
		commonComment.User.FollowerCount = commonCommentUser.FollowerCount
		isFollow, _ := model.CheckUserFollowById(userid, modelComment.FromUserId)
		commonComment.User.IsFollow = isFollow

		commonCommentList = append(commonCommentList, commonComment)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg: errmsg.OK.StatusMsg,
		},
		CommentList: commonCommentList,
	})
}

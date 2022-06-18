// Package favorite provides favorite for api.
package favorite

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)

type FavoriteListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
}

func GetFavoriteList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	// 检查用户是否存在
	userId, _ := redis.RCGet(token).Int64()
	// user喜欢的视频列表Id
	userLikeVideoList := model.GetFavoriteListById(int64(user_id))
	// 根据视频Id获取视频信息
	var modelVideoList []model.Video
	for _, userLikeVideoId := range userLikeVideoList {
		userLikeVideoInfo, _ := model.GetVideoInfo(userLikeVideoId)
		modelVideoList = append(modelVideoList, userLikeVideoInfo)
	}

	var commonVideoList []common.Video

	for _, modelVideo := range modelVideoList {
		var commonVideo common.Video
		commonVideo.Id = modelVideo.Id
		commonVideo.PlayUrl = modelVideo.PlayUrl
		commonVideo.CoverUrl = modelVideo.CoverUrl
		commonVideo.FavoriteCount = modelVideo.FavoriteCount
		isFavorite, _ := model.CheckUserLikeVideoById(userId, modelVideo.Id)
		commonVideo.IsFavorite = isFavorite
		_, count := model.GetCommentCountById(modelVideo.Id)
		commonVideo.CommentCount = count
		modelVideoAuthorInfos, _ := model.GetUserById(modelVideo.Author)
		commonVideo.Author.FollowCount = modelVideoAuthorInfos.FollowCount
		commonVideo.Author.FollowerCount = modelVideoAuthorInfos.FollowerCount
		commonVideo.Author.Id = modelVideoAuthorInfos.Id
		commonVideo.Author.Name = modelVideoAuthorInfos.Username
		isFollow, _ := model.CheckUserFollowById(userId, modelVideoAuthorInfos.Id)
		commonVideo.Author.IsFollow = isFollow
		commonVideo.Title = modelVideo.Title
		commonVideoList = append(commonVideoList, commonVideo)
	}

	c.JSON(http.StatusOK, FavoriteListResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		},
		VideoList: commonVideoList,
	})
}

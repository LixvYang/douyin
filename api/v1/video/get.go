// Package video provides video for api.
package video

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

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// GetVideoList returns video list.
// 通过视频Id查询视频信息
func GetVideoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v, err := model.GetVideoInfo(int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, "获取视频信息失败")
		return
	}
	c.JSON(http.StatusOK, &model.Video{
		Id:            v.Id,
		Author:        v.Author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CreateTime:    v.CreateTime,
	})
}

func GetVideoList(c *gin.Context) {
	lastTime, _ := strconv.Atoi(c.Query("last_time"))
	token := c.Query("token")
	userid, _ := redis.RCGet(token).Int64()
	// user, exist := user.TokenToUser[token]
	// if !exist {
	// 	fmt.Println("用户未登录")
	// }

	var code int64
	var modelVideoList []model.Video

	if lastTime != 0 {
		modelVideoList, code = model.ListVideosByCount(30, int64(lastTime))
		if code == errmsg.ERROR.StatusCode {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: errmsg.ErrDatabase.StatusCode,
				StatusMsg:  errmsg.ErrBind.StatusMsg,
			})
			return
		}
	} else {
		modelVideoList, code = model.ListVideosByCount(30, 0)
		if code == errmsg.ERROR.StatusCode {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: errmsg.ErrDatabase.StatusCode,
				StatusMsg:  errmsg.ErrBind.StatusMsg,
			})
			return
		}
	}

	var commonVideoList []common.Video
	// isFollow 和 作者信息  isFavorite
	for _, modelVideo := range modelVideoList {
		var commonVideo common.Video
		commonVideo.Id = modelVideo.Id

		// 作者信息
		modelVideoAuthorInfos, _ := model.GetUserById(modelVideo.Author)
		commonVideo.Author.FollowCount = modelVideoAuthorInfos.FollowCount
		commonVideo.Author.FollowerCount = modelVideoAuthorInfos.FollowerCount
		commonVideo.Author.Id = modelVideoAuthorInfos.Id
		commonVideo.Author.Name = modelVideoAuthorInfos.Username
		// 是否关注
		isFollow, _ := model.CheckUserFollowById(userid, modelVideoAuthorInfos.Id)
		commonVideo.Author.IsFollow = isFollow

		commonVideo.PlayUrl = modelVideo.PlayUrl
		commonVideo.CoverUrl = modelVideo.CoverUrl

		// FIXME FavoriteCount的总数计算
		commonVideo.FavoriteCount = modelVideo.FavoriteCount

		_, count := model.GetCommentCountById(modelVideo.Id)
		commonVideo.CommentCount = count

		isFavorite, _ := model.CheckUserLikeVideoById(userid, modelVideo.Id)
		commonVideo.IsFavorite = isFavorite
		commonVideo.Title = modelVideo.Title
		commonVideoList = append(commonVideoList, commonVideo)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		},
		NextTime:  time.Now().Unix(),
		VideoList: commonVideoList,
	})
}


// 用户发布列表
func GetUserPublishList(c *gin.Context)  {
		// token := c.Query("token")
		userId, _ := strconv.Atoi(c.Query("user_id"))
		// if userId !=  {
		// 	c.JSON(http.StatusOK, common.Response{
		// 		StatusCode: errmsg.ErrUserNotFound.StatusCode,
		// 		StatusMsg:  errmsg.ErrUserNotFound.StatusMsg,
		// 	})
		// 	return
		// }

		modelVideoIdList, code := model.GetVideosByAuthorId(int64(userId))
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: errmsg.ErrDatabase.StatusCode,
			})
		} 

		var commonVideoList []common.Video
		for _, modelVideoId := range modelVideoIdList {
			var commonVideo common.Video
			modelVideo, _ := model.GetVideoInfo(modelVideoId)
			commonVideo.Id = modelVideoId
			commonVideo.PlayUrl = modelVideo.PlayUrl
			commonVideo.CoverUrl = modelVideo.CoverUrl
			commonVideo.FavoriteCount = modelVideo.FavoriteCount
			_, commentCount := model.GetCommentCountById(modelVideoId)
			commonVideo.CommentCount = commentCount
			isFavorite, _ := model.CheckUserLikeVideoById(int64(userId), modelVideoId)
			commonVideo.IsFavorite = isFavorite
			author, _ := model.GetUserById(modelVideo.Author)
			commonVideo.Author.Id = author.Id
			commonVideo.Author.Name = author.Username
			commonVideo.Author.FollowCount = author.FollowCount
			commonVideo.Author.FollowerCount = author.FollowerCount
			isFollow, _ := model.CheckUserFollowById(int64(userId), author.Id)
			commonVideo.Author.IsFollow = isFollow
			commonVideo.Title = modelVideo.Title
			commonVideoList = append(commonVideoList, commonVideo)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: errmsg.OK.StatusCode,
				StatusMsg:  errmsg.OK.StatusMsg,
			},
			VideoList: commonVideoList,
		})
}
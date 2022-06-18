// Package follow provides follow for api.
package follow

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)

type FollowResponse struct {
	common.Response
	UserList []common.UserInfo `json:"user_list"`
}

func GetFollowList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	userid, _ := redis.RCGet(token).Int64()

	if int64(userid) != userid {
		c.JSON(http.StatusBadRequest, "用户不存在")
		return
	}

	modelUserList, code := model.GetUserFollowId(int64(user_id))
	if code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, "获取关注列表失败")
	}

	var UserInfoList []common.UserInfo
	for _, modelUserId := range modelUserList {
		var UserInfo common.UserInfo
		modelUser, _ := model.GetUserById(modelUserId)
		UserInfo.Id = modelUser.Id
		UserInfo.Name = modelUser.Username
		// FIXME FollowerCount和FollowCount总数计算，防止恶意修改数据库使得总数和列表数量不一致
		UserInfo.FollowCount = modelUser.FollowCount
		UserInfo.FollowerCount = modelUser.FollowerCount
		isFollow, _ := model.CheckUserFollowById(int64(user_id), modelUserId)
		UserInfo.IsFollow = isFollow
		UserInfoList = append(UserInfoList, UserInfo)
	}

	c.JSON(http.StatusOK, FollowResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		},
		UserList: UserInfoList,
	})
}

func GetFollowerList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	userid, _ := redis.RCGet(token).Int64()

	if userid != int64(user_id) {
		c.JSON(http.StatusBadRequest, "用户不存在")
		return
	}

	modelUserList, code := model.GetUserFollowerId(int64(user_id))
	if code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, "获取粉丝列表失败")
	}

	var UserInfoList []common.UserInfo
	for _, modelUserId := range modelUserList {
		var UserInfo common.UserInfo
		modelUser, _ := model.GetUserById(modelUserId)
		UserInfo.Id = modelUser.Id
		UserInfo.Name = modelUser.Username
		UserInfo.FollowCount = modelUser.FollowCount
		UserInfo.FollowerCount = modelUser.FollowerCount
		isFollow, _ := model.CheckUserFollowById(int64(user_id), modelUserId)
		UserInfo.IsFollow = isFollow
		UserInfoList = append(UserInfoList, UserInfo)
	}

	c.JSON(http.StatusOK, FollowResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		},
		UserList: UserInfoList,
	})

}

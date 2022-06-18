// Package user provides user for api.
package user

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	common.Response
	User common.UserInfo `json:"user"`
}

// 通过用户ID 请求用户信息
func GetUserById(c *gin.Context) {
	// 要查询用户的ID
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")

	// var UserInfo common.UserInfo
	data, code := model.GetUser(userId)
	if code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: common.Response{
				StatusCode: code,
				StatusMsg:  errmsg.ERROR.StatusMsg,
			},
		})
		return
	}

	UserInfo := &common.UserInfo{
		Id:            data.Id,
		Name:          data.Username,
		FollowCount:   data.FollowCount,
		FollowerCount: data.FollowerCount,
		IsFollow:      false,
	}

	if token != "" {
		userId := redis.RCGet(token).Val()
		userid, _ := strconv.Atoi(userId)
		isFollow, _ := model.CheckUserFollowById(int64(userid), data.Id)
		if isFollow {
			UserInfo.IsFollow = true
		}
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg:  errmsg.OK.StatusMsg,
		},
		User: *UserInfo,
	})
	return
}

// Package user provides user for api.
package user

import (
	"net/http"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/crypto"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	common.Response
	UserId int64    `json:"user_id"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	formData, code := model.CheckLogin(username, password)
	if code == errmsg.ERROR.StatusCode {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{
				StatusCode: code,
				StatusMsg:  errmsg.ERROR.StatusMsg,
			},
		})
		return
	}
	userid, _ := model.GetUserIdByName(username)
	token := utils.MD5WithSalt(formData.Username)
	redis.RCSet(token, userid, time.Hour)
	c.JSON(http.StatusOK, LoginResponse{
		Response: common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg: errmsg.OK.StatusMsg,
		},
		UserId: formData.Id,
		Token:  token,
	})
}

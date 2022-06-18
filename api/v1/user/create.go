// Package user provides user for api.
package user

import (
	"fmt"
	"net/http"
	"tiktok/model"
	"tiktok/pkg/common"
	"tiktok/pkg/crypto"
	"tiktok/pkg/errmsg"
	"tiktok/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateUserResponse struct {
	common.Response
	UserId     int64    `json:"user_id"`
	Token      string `json:"token"`
}

// 创建用户 通过 username, password
func AddUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, CreateUserResponse{
			Response: common.Response{
				StatusCode: errmsg.ErrUserRegister.StatusCode,
				StatusMsg:  errmsg.ErrUserRegister.StatusMsg,
			},
		})
		return
	}

	// 判断要创建的用户名是否已经存在
	code := model.CheckUser(username)
	if code == errmsg.ErrUsernameExist.StatusCode {
		c.JSON(http.StatusOK, CreateUserResponse{
			Response: common.Response{
				StatusCode: code,
				StatusMsg:  errmsg.ErrUsernameExist.StatusMsg,
			},
		})
		return
	}

	data := &model.User{
		Username: username,
		Password: password,
	}
	
	if code := model.AddUser(data); code != errmsg.OK.StatusCode {
		c.JSON(http.StatusOK, CreateUserResponse{
			Response: common.Response{
				StatusCode: code,
				StatusMsg:  errmsg.ERROR.StatusMsg,
			},
		})
		return
	}
	token := utils.MD5WithSalt(username)
	userId, code := model.GetUserIdByName(username)
	if code == errmsg.ERROR.StatusCode {
		c.JSON(http.StatusOK, CreateUserResponse{
			Response: common.Response{
				StatusCode: code,
				StatusMsg:  errmsg.ERROR.StatusMsg,
			},
		})
		return
	}
	redis.RCSet(token, userId, time.Hour)
	result := redis.RCGet(token).Val()
	fmt.Println(result)
	c.JSON(http.StatusOK, CreateUserResponse{
		Response: common.Response{StatusCode: 0},
		UserId: data.Id,
		Token: token,
	})
}

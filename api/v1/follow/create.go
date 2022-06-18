// Package follow provides follow for api.
package follow

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

func AddFollow(c *gin.Context)  {
	token := c.Query("token")
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	userid, _ := redis.RCGet(token).Int64()
	var UserFollow model.UserFollow
	if action_type == 1 {
		UserFollow.UserId = userid
		UserFollow.FollowId = int64(to_user_id)
		UserFollow.IsFollow = int64(action_type)
		UserFollow.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		code := model.AddFollow(&UserFollow)
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "关注失败")
		}

		code2 := model.SetFollowCount(userid, int64(to_user_id), int64(action_type))
		if code2 != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "关注总数、粉丝总数 + 1 失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg: errmsg.OK.StatusMsg,
		})
	} else if action_type == 2 {
		code := model.DelFollow(userid, int64(to_user_id))
		if code != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "取消关注失败")
		}

		code2 := model.SetFollowCount(userid, int64(to_user_id), int64(action_type))
		if code2 != errmsg.OK.StatusCode {
			c.JSON(http.StatusOK, "关注总数、粉丝总数 - 1 失败")
		}

		c.JSON(http.StatusOK, common.Response{
			StatusCode: errmsg.OK.StatusCode,
			StatusMsg: errmsg.OK.StatusMsg,
		})
	}
}
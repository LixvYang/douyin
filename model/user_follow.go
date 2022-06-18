// Package model.UserFollow provides the UserFollow model.
package model

import "tiktok/pkg/errmsg"

type UserFollow struct {
	Id         int64  `gorm:"id" json:"id"`                   // 自增ID
	UserId     int64  `gorm:"user_id" json:"user_id"`         // 用户ID
	FollowId   int64  `gorm:"follow_id" json:"follow_id"`     // 关注的用户ID
	IsFollow   int64  `gorm:"is_follow" json:"is_follow"`     // 是否关注 0:取消关注 1:关注
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
}

func (*UserFollow) TableName() string {
	return "user_follow"
}

// 根据user_id查询关注的用户
func GetUserFollowId(id int64) ([]int64, int64) {
	var userId []int64
	var userIdFollow []UserFollow

	// 绑定数据表
	if err := db.Model(&userIdFollow).Error; err != nil {
		return userId, errmsg.ERROR.StatusCode
	}

	if err := db.Model(&userIdFollow).Select("follow_id").Where("user_id = ? AND is_follow = ?", id, 1).Find(&userId).Error; err != nil {
		return userId, errmsg.ERROR.StatusCode
	}
	return userId, errmsg.OK.StatusCode
}

// 查询用户的粉丝 根据user_id
// select user_id from user_follow where follow_id = ?(userId)
func GetUserFollowerId(id int64) ([]int64, int64) {
	var userId []int64
	var userIdFollow []UserFollow

	// 绑定数据表
	if err := db.Model(&userIdFollow).Error; err != nil {
		return userId, errmsg.ERROR.StatusCode
	}

	if err := db.Model(&userIdFollow).Select("user_id").Where("follow_id = ? AND is_follow = ?", id, 1).Find(&userId).Error; err != nil {
		return userId, errmsg.ERROR.StatusCode
	}
	return userId, errmsg.OK.StatusCode
}

// 查看userId 是否关注了 followId
// Select is_follow from user_follow where user_id = ?(userId) and follow_id = ?(followId)
// if is_follow == 0 return false
// if is_follow == 1 return true
func CheckUserFollowById(userId, followId int64) (bool, int64) {
	var isFollow int
	var userFollow UserFollow
	if err := db.Model(&userFollow).Error; err != nil {
		return false, errmsg.ERROR.StatusCode
	}

	if err := db.Model(&userFollow).Select("is_follow").Where("user_id = ? AND follow_id = ?", userId, followId).Find(&isFollow).Error; err != nil {
		return false, errmsg.ERROR.StatusCode
	}
	if isFollow == 1 {
		return true, errmsg.OK.StatusCode
	}
	return false, errmsg.ERROR.StatusCode
}

func AddFollow(data *UserFollow) int64 {
	var userFollow UserFollow
	if err := db.Model(&userFollow).Create(&data).Error; err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}

func DelFollow(user_id int64, to_user_id int64) int64 {
	var userFollow UserFollow
	if err := db.Model(&userFollow).Where("user_id = ? AND follow_id = ?", user_id, to_user_id).Delete(&userFollow).Error; err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}

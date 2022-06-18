// Package model.UsersLikeVideo provides the UsersLikeVideo model.
package model

import "tiktok/pkg/errmsg"

type UsersLikeVideo struct {
	Id      int64 `gorm:"id" json:"id"`
	UserId  int64 `gorm:"user_id" json:"user_id"`   // 用户ID
	VideoId int64 `gorm:"video_id" json:"video_id"` // 视频ID
	IsLike  int64 `gorm:"is_like" json:"is_like"`   // 是否点赞 0:取消点赞 1:点赞
}

func (*UsersLikeVideo) TableName() string {
	return "users_like_video"
}

func CheckUserLikeVideoById(userId int64, videoId int64) (bool, int64) {
	var isLike int
	var userLikeVideo UsersLikeVideo
	if err := db.Model(&userLikeVideo).Error; err != nil {
		return false, errmsg.ERROR.StatusCode
	}

	if err := db.Model(&userLikeVideo).Select("is_like").Where("user_id = ? AND video_id = ?", userId, videoId).Find(&isLike).Error; err != nil {
		return false, errmsg.ERROR.StatusCode
	}
	if isLike == 1 {
		return true, errmsg.OK.StatusCode
	}
	return false, errmsg.ERROR.StatusCode
}

func AddFavorite(data *UsersLikeVideo) int64 {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}

func DelFavorite(userId int64, videoId int64) int64 {
	var data UsersLikeVideo
	err := db.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&data).Error
	if err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}
	
// 查询用户id下的点赞视频
// select video_id from users_like_video where user_id = ?
func GetFavoriteListById(userId int64) []int64 {
	var userLikeVideo UsersLikeVideo
	var data []int64
	err := db.Model(&userLikeVideo).Select("video_id").Where("user_id = ?", userId).Find(&data).Error; if err != nil {
		return nil
	} 
	return data
}
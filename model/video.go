// Package model.video provides video model.
package model

import (
	"gorm.io/gorm"
	"tiktok/pkg/errmsg"
)

type Video struct {
	Id            int64    `gorm:"id" json:"id"`
	Author        int64    `gorm:"author" json:"author"`                 // 作者id
	PlayUrl       string `gorm:"play_url" json:"play_url"`             // 播放地址
	CoverUrl      string `gorm:"cover_url" json:"cover_url"`           // 封面地址
	FavoriteCount int64    `gorm:"favorite_count" json:"favorite_count"` // 喜欢数
	CommentCount  int64    `gorm:"comment_count" json:"comment_count"`   // 评论数
	CreateTime    int64 `gorm:"create_time" json:"create_time"`       // 创建时间
	Title		  string  	`gorm:"title" json:"title"`                 // 视频标题
}

func (*Video) TableName() string {
	return "video"
}

func GetVideoInfo(id int64) (Video, error) {
	var video Video
	err := db.Where("id = ?", id).Find(&video).Error
	if err != nil {
		return video, err
	}
	return video, err
}

func ListVideos() ([]Video, error) {
	video := []Video{}
	videos := make([]Video, 0)

	if err := db.Model(&video).Error; err != nil {
		return videos, err
	}

	if err := db.Model(&video).Where("").Order("id desc").Find(&videos).Error; err != nil {
		return videos, err
	}
	return videos, err
}

// 通过count控制视频数量
func ListVideosByCount(MaxCount int, lastTime int64) ([]Video, int64) {
	videos := make([]Video, 0)
	if err := db.Model(&videos).Where("create_time > ?", lastTime).Order("id desc").Limit(MaxCount).Find(&videos).Error; err != nil {
		return videos, errmsg.ERROR.StatusCode
	}
	return videos, errmsg.OK.StatusCode
}

func AddVideo(data *Video) int64 {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}
	
// 根据视频id 查询作者信息
func GetAuthorInfo(id int) (User, int64) {
	var user User
	// select author from video where id = ?
	err := db.Limit(1).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR.StatusCode
	}
	return user, errmsg.OK.StatusCode
}

// select video_id from video where author = ?
func GetVideosByAuthorId(author int64) ([]int64, int64) {
		var data []int64
		var video Video
		err := db.Model(&video).Select("id").Where("author = ?", author).Find(&data).Error
		if err != nil {
			return data, errmsg.ERROR.StatusCode
		}
		return data, errmsg.OK.StatusCode
}

// 修改点赞总数
func SetFavoriteCount(VideoId int64, actionType int64) int64 {
	video := Video{}
	video.Id = VideoId
	if actionType == 1 {
		if err := db.Model(&video).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
	}else {
		if err := db.Model(&video).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
	}
	return errmsg.OK.StatusCode
}

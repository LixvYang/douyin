// Package model.comment provides comment model.
package model

import "tiktok/pkg/errmsg"

type Comment struct {
	Id              int64    `gorm:"id" json:"id"`                               // 自增ID
	VideoId         int64    `gorm:"video_id" json:"video_id"`                   // 视频ID
	FromUserId      int64    `gorm:"from_user_id" json:"from_user_id"`           // 留言者ID
	Comment         string 	`gorm:"comment" json:"comment"`                     // 评论内容
	CreateDate      string 	`gorm:"create_date" json:"create_date"`             // 评论时间
}

func (*Comment) TableName() string {
	return "comment"
}

// FIXME
func GetCommentCountById(video_id int64) (int64, int64) {
	var comment Comment
	var count int64
	err := db.Model(&comment).Where("video_id = ?", video_id).Count(&count).Error; if err != nil {
		return errmsg.ERROR.StatusCode, count
	}
	return errmsg.OK.StatusCode, count
}

func AddComment(data *Comment) int64 {
	err := db.Create(&data).Error; if err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}

func DelComment(id int64) int64 {
	var comment Comment
	err := db.Model(&comment).Where("id = ?", id).Delete(&comment).Error; if err != nil {
		return errmsg.ERROR.StatusCode
	}
	return errmsg.OK.StatusCode
}

// select comment 
func GetCommentListByVideoId(videoId int64) ([]Comment, int64) {
	var comment Comment
	var data []Comment
	err := db.Model(&comment).Where("video_id = ?", videoId).Find(&data).Error; if err != nil {
		return nil, errmsg.ERROR.StatusCode
	}
	return data, errmsg.OK.StatusCode
}
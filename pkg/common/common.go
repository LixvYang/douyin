// Package common provides common package.
package common

// Golang 的 “omitempty” 关键字略解 https://www.jianshu.com/p/a2ed0d23d1b0
type Response struct {
	StatusCode int64    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id"`
	Author        UserInfo   `json:"author"`
	Title		  string  	`json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type Comment struct {
	Id         int64    `json:"id"`
	User       UserInfo   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type UserInfo struct {
	Id            int64    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64    `json:"follow_count"`
	FollowerCount int64    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

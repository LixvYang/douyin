// Package model.user provides the user model
package model

import (
	"gorm.io/gorm"
	"tiktok/pkg/errmsg"
)

// "time"

type User struct {
	Id            int64    `gorm:"id" json:"id"`                         // 用户ID
	Username      string `gorm:"username" json:"username"`             // 用户名称
	Password      string `gorm:"password" json:"password"`             // 用户密码
	FollowerCount int64    `gorm:"follower_count" json:"follower_count"` // 用户的粉丝总数
	FollowCount   int64    `gorm:"follow_count" json:"follow_count"`     // 用户的关注数量
}

func (u *User) TableName() string {
	return "user"
}

// 查询用户名是否存在
func CheckUser(name string) int64 {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.Id > 0 {
		return errmsg.ErrUsernameExist.StatusCode
	}
	// 用户名不存在
	return errmsg.OK.StatusCode
}

// 新增用户
func AddUser(data *User) int64 {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR.StatusCode
	}	
	return errmsg.OK.StatusCode
}

//查询用户
func GetUser(id int) (User, int64) {
	var user User
	err := db.Limit(1).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR.StatusCode
	}
	return user, errmsg.OK.StatusCode
}

func GetUserIdByName(name string) (int64,  int64) {
	var user User
	err := db.Limit(1).Where("username = ?", name).Find(&user).Error
	if err != nil {
		return user.Id, errmsg.ERROR.StatusCode
	}
	return user.Id, errmsg.OK.StatusCode
}

// 查询用户列表
func ListUsers() ([]User, int64) {
	users := make([]User, 0)
	user := User{}
	if err := db.Model(&user).Error; err != nil {
		return users, errmsg.ERROR.StatusCode
	}
	if err := db.Where("").Order("id desc").Find(&users).Error; err != nil {
		return users, errmsg.ERROR.StatusCode
	}
	return users, errmsg.OK.StatusCode
}

func GetUserById(id int64) (User, error) {
	var user User
	if err := db.Limit(1).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}


// 后台验证登录
func CheckLogin(username string, password string) (User, int64) {
	var user User
	
	db.Where("username = ?", username).First(&user)

	if user.Id == 0 {
		return user, errmsg.ERROR.StatusCode
	}
	if user.Password != password {
		return user, errmsg.ERROR.StatusCode
	}

	return user, errmsg.OK.StatusCode
}

// 修改关注总数、粉丝总数
func SetFollowCount(userId int64, toUserId int64, actionType int64) int64 {
	user1 := User{}
	user1.Id = userId
	user2 := User{}
	user2.Id = toUserId
	if actionType == 1 {
		if err := db.Model(&user1).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
		if err := db.Model(&user2).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
	}else {
		if err := db.Model(&user1).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
		if err := db.Model(&user2).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return errmsg.ERROR.StatusCode
		}
	}
	return errmsg.OK.StatusCode
}




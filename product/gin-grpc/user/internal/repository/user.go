package repository

import (
	"errors"
	"user/internal/service"

	"gorm.io/gorm"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	passwordDigest string
}

const (
	passwordDigest = 12 //密码加密难度
)

// CheckUesrExist 检查用户是否存在
func (user *User) CheckUesrExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name = ?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo 获取用户信息
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUesrExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

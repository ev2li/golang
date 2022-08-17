package repository

import (
	"errors"
	"user/internal/service"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	passwordDigest string
}

const (
	PassWordCost = 12 //密码加密难度
)

// CheckUesrExist 检查用户是否存在
func (user *User) CheckUesrExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name = ?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo 获取用户信息
func (user *User) ShowUserInfo(req *service.UserRequest) error {
	if exist := user.CheckUesrExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

// BuildUser 序列化User
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}
	return &userModel
}

// UserCreate 用户生成
func (user *User) UserCreate(req *service.UserRequest) error {
	var count int64
	DB.Where("user_name = ?", req.UserName).Count(&count)
	if count > 0 {
		return errors.New("UserName Exist")
	}

	user = &User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	_ = user.SetPassword(req.Password)
	err := DB.Create(user).Error
	return err
}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return nil
	}
	user.passwordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.passwordDigest), []byte(password)) == nil
}

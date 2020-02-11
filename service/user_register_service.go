package service

import (
	"frescape/model"
	"frescape/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=1,max=16"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=32"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=32"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 40201,
			Msg:    "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40202,
			Msg:    "用户名被占用",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		UserName: service.UserName,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: 50201,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 50202,
			Msg:    "注册失败",
		}
	}

	return user, nil
}

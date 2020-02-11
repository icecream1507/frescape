package serializer

import (
	"frescape/model"
	"time"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	//Status    string `json:"status"`
	//Avatar    string `json:"avatar"`
	CreatedAt time.Time  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		//Status:    user.Status,
		//Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}

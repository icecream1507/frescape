package model

import (
	"github.com/jinzhu/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	TopicID    uint
	ParentID   uint
	CreatorID  uint
	Content    string
	Permission uint
}

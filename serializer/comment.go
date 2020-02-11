package serializer

import (
	"frescape/model"
	"time"
)

// Comment 评论序列化器
type Comment struct {
	ID         uint
	TopicID    uint
	ParentID   uint
	CreatorID  uint
	Content    string
	Permission uint
	CreatedAt  time.Time
	Children   []model.Comment
}

// BuildComment 序列化评论
func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:         item.ID,
		TopicID:    item.TopicID,
		ParentID:   item.ParentID,
		CreatorID:  item.CreatorID,
		Content:    item.Content,
		Permission: item.Permission,
		CreatedAt:  item.CreatedAt,
	}
}

// BuildComments 序列化影像列表
func BuildComments(items []model.Comment, operatorID uint) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		if comment.Permission == 1 && operatorID != comment.CreatorID {
			comments = append(comments, Comment{
				ID: comment.ID,
				Permission: item.Permission,
			})
		} else {
			comments = append(comments, comment)
		}
	}
	return comments
}

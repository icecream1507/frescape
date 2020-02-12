package service

import (
	"frescape/model"
	"frescape/serializer"
)

// PostCommentService 评论提交的服务
type PostCommentService struct {
	TopicID    uint `form:"tid" json:"tid" binding:"required"`
	ParentID   uint `form:"pid" json:"pid"`
	CreatorID  uint
	Content    string `form:"content" json:"content" binding:"required"`
	Permission uint `form:"permission" json:"permission" binding:"required"`
}


// Post 提交评论
func (service *PostCommentService) Post() serializer.Response {
	comment := model.Comment{
		TopicID: service.TopicID,
		ParentID: service.ParentID,
		CreatorID: service.CreatorID,
		Content: service.Content,
		Permission: service.Permission,
	}

	err := model.DB.Create(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50301,
			Msg:    "评论保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildComment(comment),
		Msg: "评论提交成功",
	}

}

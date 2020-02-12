package service

import (
	"frescape/model"
	"frescape/serializer"
	"strconv"
)

type ShowCommentService struct {}

// typeProcess 类型转换
func typeProccess(str string) uint {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return uint(num)
}

// Show 根据传递过来的主题ID获取下面所有的评论
func (service *ShowCommentService) Show(tid string) serializer.Response {
	// 初始化接收数据库的变量
	var ParentComments []model.Comment

	// 获取当前主题下的所有根评论
	if err := model.DB.Where("parent_id = ? and topic_id = ?", 0, typeProccess(tid)).Find(&ParentComments).Error; err != nil {
		return serializer.Response{
			Status: 50311,
			Msg:    "读取评论失败",
			Error:  err.Error(),
		}
	}

	// 获取每一条根评论下面的所有子评论
	var CommentsResponse []serializer.Comment
	for _, ParentComment := range ParentComments {
		var ChildComments []model.Comment
		var CommentResponse serializer.Comment

		if err := model.DB.Where("parent_id = ? and topic_id = ?", ParentComment.ID, typeProccess(tid)).Find(&ChildComments).Error; err != nil {
			return serializer.Response{
				Status: 50311,
				Msg:   "读取评论失败",
				Error: err.Error(),
			}
		}

		CommentResponse = serializer.BuildComment(ParentComment)
		CommentResponse.Children = ChildComments
		CommentsResponse = append(CommentsResponse, CommentResponse)
	}

	return serializer.BuildListResponse(CommentsResponse, uint(len(CommentsResponse)), "读取评论成功")
}

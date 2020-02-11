package service

import (
	"frescape/model"
	"frescape/serializer"
)

// ListPictureService 影像列表服务
type ListPictureService struct {
	Limit 		int `form:"limit"`
	Start 		int `form:"start"`
}

// List 影像列表
func (service *ListPictureService) List(operatorID  uint) serializer.Response {
	pictures := []model.Picture{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Picture{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50131,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&pictures).Error; err != nil {
		return serializer.Response{
			Status: 50131,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildPictures(pictures, operatorID), uint(total), "获取影像列表成功")
}

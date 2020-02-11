package service

import (
	"frescape/model"
	"frescape/serializer"
)

type ShowPictureService struct {}

// Show 读取影像
func (service *ShowPictureService) Show(pictureID string, opreatorID uint) serializer.Response {
	// 读取影像
	var picture model.Picture
	err := model.DB.First(&picture, pictureID).Error
	if err != nil {
		return serializer.Response{
			Status: 40111,
			Msg:    "影像不存在",
			Error:  err.Error(),
		}
	}

	if picture.Permission == 1 && opreatorID != picture.CreatorID {
		return serializer.Response{
			Status: 40112,
			Msg:    "用户无权限",
		}
	}

	return serializer.Response{
		Data: serializer.BuildPicture(picture),
		Msg:  "影像读取成功",
	}
}

package service

import (
	"frescape/model"
	"frescape/serializer"
)

type DeletePictureService struct {}

// Delete 删除影像
func (service *DeletePictureService) Delete(pictureID string, opreatorID uint) serializer.Response {
	var picture model.Picture
	err := model.DB.First(&picture, pictureID).Error
	if err != nil {
		return serializer.Response{
			Status: 40131,
			Msg:    "影像不存在",
			Error:  err.Error(),
		}
	}

	if picture.CreatorID != opreatorID {
		return serializer.Response {
			Status: 40133,
			Msg: "用户无权限",
		}
	}

	err = model.DB.Delete(&picture).Error
	if err != nil {
		return serializer.Response{
			Status: 50121,
			Msg:    "影像删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg:    "影像删除成功",
	}
}

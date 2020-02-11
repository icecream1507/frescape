package service

import (
	"frescape/model"
	"frescape/serializer"
)

// UpdatePictureService 影像更新的服务
type UpdatePictureService struct {
	Title 		string `form:"title" json:"title" binding:"required,min=1,max=50"`
	Info  		string `form:"info" json:"info" binding:"max=3000"`
	OpreatorID 	uint
}

// Update 更新影像
func (service *UpdatePictureService) Update(pictureID string) serializer.Response {
	var picture model.Picture
	err := model.DB.First(&picture, pictureID).Error
	if err != nil {
		return serializer.Response{
			Status: 40121,
			Msg:    "影像不存在",
			Error:  err.Error(),
		}
	}

	if picture.CreatorID != service.OpreatorID {
		return serializer.Response {
			Status: 40123,
			Msg: "用户无权限",
		}
	}

	picture.Title = service.Title
	picture.Info = service.Info
	err = model.DB.Save(&picture).Error
	if err != nil {
		return serializer.Response{
			Status: 50111,
			Msg:    "影像更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data:   serializer.BuildPicture(picture),
		Msg:    "影像更新成功",
	}

}

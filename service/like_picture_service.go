package service

import (
	"frescape/model"
	"frescape/cache"
	"frescape/serializer"

	"strconv"
)

type LikePictureService struct {}

// Like 点赞影像
func (service *LikePictureService) Like(pictureID string, opreatorID uint) serializer.Response {
	var picture model.Picture
	err := model.DB.First(&picture, pictureID).Error
	if err != nil {
		return serializer.Response{
			Status: 40141,
			Msg:    "影像不存在",
			Error:  err.Error(),
		}
	}

	var msg string
	if cache.RedisClient.SIsMember("like:picture:" + pictureID, opreatorID).Val() {
		cache.RedisClient.SRem("like:picture:" + pictureID, opreatorID)
		msg = "取消点赞成功"
	} else {
		cache.RedisClient.SAdd("like:picture:" + pictureID, opreatorID)
		msg = "点赞成功"
	}

	return serializer.Response{
		Data: cache.RedisClient.SMembers("like:picture:" + strconv.Itoa(int(picture.ID))).Val(),
		Msg:  msg,
	}
}

package serializer

import (
	"frescape/cache"
	"frescape/model"
	"time"
	"strconv"
)

// Picture 影像序列化器
type Picture struct {
	ID    		uint   `json:"id"`
	Title 		string `json:"title"`
	Info  		string `json:"info"`
	Like		int64 `json:"like"`
	CreatorID 	uint `json:"creator_id"`
	Permission  uint `json:"permission"`
	Key   		string `json:"key"`
	URL   		string `json:"url"`
	//View      uint64 `json:"view"`
	CreatedAt 	time.Time `json:"created_at"`
}

// BuildPicture 序列化影像
func BuildPicture(item model.Picture) Picture {
	return Picture{
		ID:    item.ID,
		Title: item.Title,
		Info:  item.Info,
		Like:  cache.RedisClient.SCard("like:picture:" + strconv.Itoa(int(item.ID))).Val(),
		CreatorID: item.CreatorID,
		Permission: item.Permission,
		Key:   item.Key,
		URL:   item.SignGetURL(),
		//View:      item.View(),
		CreatedAt: item.CreatedAt,
	}
}

// BuildPictures 序列化影像列表
func BuildPictures(items []model.Picture, operatorID uint) (pictures []Picture) {
	for _, item := range items {
		picture := BuildPicture(item)
		if picture.Permission == 1 && operatorID != picture.CreatorID {
			pictures = append(pictures, Picture{
				ID: picture.ID,
				Permission: item.Permission,
			})
		} else {
			pictures = append(pictures, picture)
		}
	}
	return pictures
}

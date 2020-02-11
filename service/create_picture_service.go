package service

import (
	"frescape/cache"
	"frescape/model"
	"frescape/serializer"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// CreatePictureService 影像创建的服务
type CreatePictureService struct {
	Title      string `form:"title" json:"title" binding:"required,min=1,max=64"`
	Info       string `form:"info" json:"info" binding:"max=3000"`
	Filename   string `form:"filename" json:"filename"`
	CreatorID  uint
	Permission uint `form:"permission" json:"permission" binding:"required"`
}

// Create 创建影像
func (service *CreatePictureService) Create() serializer.Response {
	// 获取扩展名
	ext := filepath.Ext(service.Filename)

	//检测内容是否为图片
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return serializer.Response{
			Status: 40101,
			Msg:    "格式不支持",
		}
	}

	// 创建oss对象
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50101,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50101,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "pictures/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传URL
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50101,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	picture := model.Picture{
		Title:      service.Title,
		Info:       service.Info,
		CreatorID:  service.CreatorID,
		Permission: service.Permission,
		Key:        key,
	}

	err = model.DB.Create(&picture).Error
	if err != nil {
		return serializer.Response{
			Status: 50102,
			Msg:    "影像保存失败",
			Error:  err.Error(),
		}
	}

	cache.RedisClient.SAdd("like:picture"+ strconv.Itoa(int(picture.ID)), 0)
	cache.RedisClient.SRem("like:picture"+ strconv.Itoa(int(picture.ID)), 0)

	return serializer.Response{
		Data: map[string]interface{}{
			"picture":      serializer.BuildPicture(picture),
			"signedPutURL": signedPutURL,
		},
		Msg: "影像创建成功",
	}

}

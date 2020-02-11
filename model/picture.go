package model

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Picture 图片模型
type Picture struct {
	gorm.Model
	Title 		string
	Info  		string
	CreatorID	uint
	Permission  uint
	Key   		string
}
// SignGetURL 返回签名影像URL
func (picture *Picture) SignGetURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(picture.Key, oss.HTTPGet, 600)
	return signedGetURL
}
                  
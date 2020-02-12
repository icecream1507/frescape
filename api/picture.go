package api

import (
	"frescape/service"
	"frescape/serializer"
	"github.com/labstack/echo"
)

// CreatePicture 创建影像
func CreatePicture(c echo.Context) error {
	if user := CurrentUser(c); user != nil {
		service := service.CreatePictureService{CreatorID: user.ID,}
		if err := c.Bind(&service); err == nil {
			res := service.Create()
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	}
	return c.JSON(200, serializer.Response{
		Status: 40102,
		Msg: "用户未登录",
	})
}

// ShowPicture 读取影像
func ShowPicture(c echo.Context) error {
	var res serializer.Response
	var service service.ShowPictureService
	if user := CurrentUser(c); user != nil {
		res = service.Show(c.Param("id"), user.ID)
	} else {
		res = service.Show(c.Param("id"), 0)
	}
	return c.JSON(200, res)
}

// UpdatePicture 更新影像
func UpdatePicture(c echo.Context) error {
	if user := CurrentUser(c); user != nil {
		service := service.UpdatePictureService{OpreatorID: user.ID,}
		if err := c.Bind(&service); err == nil {
			res := service.Update(c.Param("id"))
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	}
	return c.JSON(200, serializer.Response{
		Status: 40122,
		Msg: "用户未登录",
	})
}

// DeletePicture 删除影像
func DeletePicture(c echo.Context) error{
	if user := CurrentUser(c); user != nil {
		var service service.DeletePictureService
		res := service.Delete(c.Param("id"), user.ID)
		return c.JSON(200, res)
	}
	return c.JSON(200, serializer.Response{
		Status: 40132,
		Msg: "用户未登录",
	})
}

// ListPicture 影像列表
func ListPicture(c echo.Context) error {
	var service service.ListPictureService
	if user := CurrentUser(c); user != nil {
		if err := c.Bind(&service); err == nil {
			res := service.List(user.ID)
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	} else {
		if err := c.Bind(&service); err == nil {
			res := service.List(0)
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	}
}

// LikePicture 影像点赞
func LikePicture(c echo.Context) error {
	if user := CurrentUser(c); user != nil {
		var service service.LikePictureService
		res := service.Like(c.Param("id"), user.ID)
		return c.JSON(200, res)
	}
	return c.JSON(200, serializer.Response{
		Status: 40142,
		Msg: "用户未登录",
	})
}

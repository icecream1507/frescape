package api

import (
	"frescape/service"
	"frescape/serializer"
	"github.com/labstack/echo"
)

// PostComment 提交评论
func PostComment(c echo.Context) error {
	if user := CurrentUser(c); user != nil {
		service := service.PostCommentService{CreatorID: user.ID,}
		if err := c.Bind(&service); err == nil {
			res := service.Post()
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	}
	return c.JSON(200, serializer.Response{
		Status: 40301,
		Msg: "用户未登录",
	})
}

// ShowComment 读取评论
func ShowComment(c echo.Context) error {
	service := service.ShowCommentService{}
	if user := CurrentUser(c); user != nil {
		if err := c.Bind(&service); err == nil {
			res := service.Show(c.Param("id"))
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	} else {
		if err := c.Bind(&service); err == nil {
			res := service.Show(c.Param("id"))
			return c.JSON(200, res)
		} else {
			return c.JSON(200, err)
		}
	}
	
}

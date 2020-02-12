package api

import (
	"frescape/model"
	"frescape/serializer"
	"frescape/service"

	"github.com/labstack/echo"
	//"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

// CurrentUser 获取当前用户
func CurrentUser(c echo.Context) *model.User {
	if user := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// UserRegister 用户注册
func UserRegister(c echo.Context) error {
	var service service.UserRegisterService
	if err := c.Bind(&service); err == nil {
		res := service.Register()
		return c.JSON(200, res)
	} else {
		return c.JSON(200, err)
	}
}

// UserLogin 用户登录
func UserLogin(c echo.Context) error {
	var service service.UserLoginService
	if err := c.Bind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			return c.JSON(200, err)
		} else {
			// 设置Session
			s, _ := session.Get("user", c)
			s.Values["user_id"] = user.ID
			s.Save(c.Request(), c.Response())
			res := serializer.Response{
				Data: user,
			}
			return c.JSON(200, res)
		}
	} else {
		return c.JSON(200, err)
	}
}

// UserMe 用户详情
func UserMe(c echo.Context) error {
	user := CurrentUser(c)
	if user != nil{
		res := serializer.Response{
			Data: user,
		}
		return c.JSON(200, res)
	} else {
		res := serializer.Response {
			Status: 40221,
			Msg: "用户未登录",
		}
		return c.JSON(200, res)
	}
}

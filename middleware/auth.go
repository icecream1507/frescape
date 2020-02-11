package middleware

import (
	"frescape/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// CurrentUser 获取登录用户
func CurrentUser(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("user", c)
		uid, ok := sess.Values["user_id"]
		if ok {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		return h(c)
	}
}

// UserAuth 登录状态验证
func UserAuth(c echo.Context) bool {
	if user := c.Get("user"); user != nil {
		if _, ok := user.(*model.User); ok {
			return true
		}
	}
	return false
}

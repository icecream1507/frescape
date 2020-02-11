package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// Session 初始化session
func Session(secret string) echo.MiddlewareFunc {
	store := sessions.NewCookieStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	*store.Options = sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}
	return session.Middleware(store)
}

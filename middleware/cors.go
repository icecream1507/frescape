package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Cors 跨域中间件
func Cors() echo.MiddlewareFunc {
	config := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Cookie"},
		AllowCredentials: true,
	}
	return middleware.CORSWithConfig(config)
}

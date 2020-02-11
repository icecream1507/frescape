package main

import (
	"frescape/cache"
	"frescape/api"
	"frescape/middleware"
	"frescape/model"
	"os"

	"github.com/labstack/echo"
	"github.com/joho/godotenv"
)

func main() {
	// 初始化
	godotenv.Load()
	model.Init(os.Getenv("MYSQL_DSN"))
	cache.Init()

	e := echo.New()
	
	// 使用中间件
	e.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	e.Use(middleware.Cors())
	e.Use(middleware.CurrentUser)

	// 用户操作
	e.POST("/api/user/register", api.UserRegister)
	e.POST("/api/user/login", api.UserLogin)
	e.GET("/api/user/me", api.UserMe)
	//e.DELETE("/api/user/logout", api.UserLogout)

	// 影像操作
	e.POST("/api/picture/upload", api.CreatePicture)
	e.GET("/api/picture/show/:id", api.ShowPicture)
	e.PUT("/api/picture/update/:id", api.UpdatePicture)
	e.DELETE("/api/picture/delete/:id", api.DeletePicture)
	e.GET("/api/picture/list", api.ListPicture)
	e.GET("/api/picture/like/:id", api.LikePicture)

	// 评论操作
	e.POST("/api/comment/post", api.PostComment)
	e.GET("/api/comment/show/:id", api.ShowComment)

	e.Logger.Fatal(e.Start(":1323"))
}

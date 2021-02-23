package routers

import (
	"api_server/internal/controllers"
	"api_server/internal/middleware"
	"api_server/store"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Init() *gin.Engine {
	r := gin.Default()
	config := new(store.Config)
	gin.SetMode(config.Get().Runmode)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "*",
		ExposedHeaders:  "*",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.GET("health", func(context *gin.Context) {
		context.JSON(200, map[string]string{"msg": "ok"})
		return
	})

	userController := new(controllers.UserController)
	publicController := new(controllers.PublicController)
	openApi := r.Group("open")
	openApi.POST("/user/login", userController.Login)
	openApi.GET("/public/short-url", publicController.GetShortUrl)
	openApi.GET("/public/js-sdk", publicController.GetJssdk)

	authApi := r.Group("auth").Use(middleware.VerifyToken())

	mallController := new(controllers.MallController)
	authApi.GET("goods/info", mallController.GetGoodsInfo)
	authApi.POST("order/create", mallController.CreateOrder)
	return r
}

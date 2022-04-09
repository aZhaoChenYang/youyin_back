package router

import (
	"github.com/gin-gonic/gin"
	"youyin/controller"
	"youyin/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1.0")
		{
			v1.POST("/login", controller.Login)
			v1.POST("/upload", middleware.JwtAuthMiddleware(), controller.Upload)
			//v1.POST("/upload", controller.Upload)
			// 管理员路由
			admin := v1.Group("/admin", middleware.JwtAuthMiddleware())
			{
				admin.POST("", controller.AddAdmin)
				admin.GET("", controller.GetAllAdmin)
				admin.DELETE("", controller.DeleteAdmin)
				admin.PUT("", controller.UpdateAdmin)
			}
			// 门店路由
			shop := v1.Group("/shop", middleware.JwtAuthMiddleware())
			{
				shop.POST("", controller.AddShop)
				shop.GET("", controller.GetShopList)
				shop.PUT("", controller.UpdateShop)
				shop.DELETE("", controller.DeleteShop)
			}
			tag := v1.Group("/tag", middleware.JwtAuthMiddleware())
			{
				tag.POST("", controller.AddTag)
				tag.GET("", controller.GetTagList)
				tag.DELETE("", controller.DeleteTag)
			}
			people := v1.Group("/people", middleware.JwtAuthMiddleware())
			{
				people.POST("", controller.AddPeople)
				people.GET("", controller.GetPeopleList)
				people.DELETE("", controller.DeletePeople)
			}
			swiper := v1.Group("/swiper", middleware.JwtAuthMiddleware())
			{
				swiper.POST("", controller.AddSwiper)
				swiper.GET("", controller.GetSwiperList)
				swiper.PUT("", controller.UpdateSwiper)
				swiper.DELETE("", controller.DeleteSwiper)
			}
			setting := v1.Group("/setting", middleware.JwtAuthMiddleware())
			{
				setting.GET("", controller.GetSetting)
				setting.PUT("", controller.UpdateSetting)
			}
			script := v1.Group("/script", middleware.JwtAuthMiddleware())
			{
				script.POST("", controller.AddScript)
				script.GET("", controller.GetScriptList)
				script.GET("/get", controller.GetScript)

				script.PUT("", controller.UpdateScript)
				script.DELETE("", controller.DeleteScript)
			}

		}

	}

	app := r.Group("/app")
	{
		app.GET("/swiper", controller.GetSwiperList)
		app.GET("/script", controller.GetScriptList)
		app.GET("/setting", controller.GetSetting)
		app.GET("/shop", controller.GetShopList)
		app.GET("/tag", controller.GetTagList)
		app.GET("/people", controller.GetPeopleList)
	}

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	App "youyin/controller/app"
	"youyin/controller/web"
	"youyin/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1.0")
		{
			v1.POST("/login", web.Login)
			v1.POST("/upload", middleware.JwtAuthMiddleware(), web.Upload)
			v1.GET("/index", middleware.JwtAuthMiddleware(), web.GetIndex)
			v1.POST("/getUserInfo", middleware.JwtAuthMiddleware(), web.GetAdminInfo)
			//v1.POST("/upload", controller.Upload)
			// 管理员路由
			admin := v1.Group("/admin", middleware.JwtAuthMiddleware())
			{
				admin.POST("", web.AddAdmin)
				admin.GET("", web.GetAllAdmin)
				admin.DELETE("", web.DeleteAdmin)
				admin.PUT("", web.UpdateAdmin)
			}
			// 门店路由
			shop := v1.Group("/shop", middleware.JwtAuthMiddleware())
			{
				shop.POST("", web.AddShop)
				shop.GET("", web.GetShopList)
				shop.PUT("", web.UpdateShop)
				shop.DELETE("", web.DeleteShop)
				shop.GET("/get", web.GetShopByID)
			}
			tag := v1.Group("/tag", middleware.JwtAuthMiddleware())
			{
				tag.POST("", web.AddTag)
				tag.GET("", web.GetTagList)
				tag.DELETE("", web.DeleteTag)
			}
			people := v1.Group("/people", middleware.JwtAuthMiddleware())
			{
				people.POST("", web.AddPeople)
				people.GET("", web.GetPeopleList)
				people.DELETE("", web.DeletePeople)
			}
			swiper := v1.Group("/swiper", middleware.JwtAuthMiddleware())
			{
				swiper.POST("", web.AddSwiper)
				swiper.GET("", web.GetSwiperList)
				swiper.PUT("", web.UpdateSwiper)
				swiper.DELETE("", web.DeleteSwiper)
			}
			setting := v1.Group("/setting", middleware.JwtAuthMiddleware())
			{
				setting.GET("", web.GetSetting)
				setting.PUT("", web.UpdateSetting)
			}
			script := v1.Group("/script", middleware.JwtAuthMiddleware())
			{
				script.POST("", web.AddScript)
				script.GET("", web.GetScriptList)
				script.GET("/get", web.GetScript)

				script.PUT("", web.UpdateScript)
				script.DELETE("", web.DeleteScript)
			}
			user := v1.Group("/user", middleware.JwtAuthMiddleware())
			{
				user.GET("", web.GetUserList)
				user.PUT("", web.UpdateUser)
			}
			order := v1.Group("/order", middleware.JwtAuthMiddleware())
			{
				order.POST("", web.AddOrder)
				order.GET("", web.GetOrderList)
				//order.PUT("", web.UpdateOrder)
			}

		}

	}

	app := r.Group("/app")
	{
		v1_0 := app.Group("/v1.0")
		{
			v1_0.GET("/swiper", web.GetSwiperList)
			v1_0.POST("/login", App.Login)
			v1_0.GET("/shop", App.GetShopList)
			v1_0.POST("/user/phone", middleware.JwtAuthMiddleware(), App.UpdateUserPhone)
		}

	}

	return r
}

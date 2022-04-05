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
			// 管理员路由
			admin := v1.Group("/admin", middleware.JwtAuthMiddleware())
			{
				admin.POST("", controller.AddAdmin)
				admin.GET("", controller.GetAdmin)
				admin.DELETE("", controller.DeleteAdmin)
				admin.PUT("", controller.UpdateAdmin)
			}
			//// 门店路由
			//shop := v1.Group("/shop", middleware.JwtAuthMiddleware())
			//{
			//
			//}
			tag := v1.Group("/tag", middleware.JwtAuthMiddleware())
			{
				tag.POST("", controller.AddTag)
				tag.GET("", controller.GetTag)
				tag.DELETE("", controller.DeleteTag)
			}
			people := v1.Group("/people", middleware.JwtAuthMiddleware())
			{
				people.POST("", controller.AddPeople)
				people.GET("", controller.GetPeople)
				people.DELETE("", controller.DeletePeople)
			}
		}

	}

	return r
}

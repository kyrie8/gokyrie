package router

import (
	"gokyrie/api"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/list", userApi.GetUserList)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
		}
	})
}

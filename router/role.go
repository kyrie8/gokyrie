package router

import (
	"gokyrie/api"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter() {
	RegistRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		roleApi := api.NewRoleApi()
		rgAuthMenu := rgAuth.Group("role")
		{
			rgAuthMenu.POST("", roleApi.AddRole)
			rgAuthMenu.PUT("/:id", roleApi.UpdateRole)
			rgAuthMenu.GET("/list", roleApi.GetRoleList)
			rgAuthMenu.PUT("/roleMenu/:id", roleApi.UpdateRoleMenu)
			rgAuthMenu.GET("/roleMenu/:id", roleApi.GetMenuByRoleId)
		}
	})
}

package router

import (
	"gokyrie/api"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter() {
	RegistRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		menuApi := api.NewMenuApi()
		rgAuthMenu := rgAuth.Group("menu")
		{
			rgAuthMenu.POST("", menuApi.AddMenu)
			rgAuthMenu.PUT("/:id", menuApi.UpdateMenu)
			rgAuthMenu.GET("/list", menuApi.GetMenuList)
		}
	})
}

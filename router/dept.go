package router

import (
	"gokyrie/api"

	"github.com/gin-gonic/gin"
)

func InitDeptRouter() {
	RegistRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		deptApi := api.NewDeptApi()
		rgAuthDept := rgAuth.Group("dept")
		{
			rgAuthDept.POST("", deptApi.AddDept)
			rgAuthDept.PUT("/:id", deptApi.UpdateDept)
			rgAuthDept.DELETE("/:id", deptApi.DeleteDeptById)
			rgAuthDept.GET("/list", deptApi.GetDeptList)
		}
	})
}

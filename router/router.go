package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "gokyrie/docs"
	"gokyrie/global"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouter []IFnRegistRoute
)

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRouter = append(gfnRouter, fn)
}

func InitRouter() {

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.New()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Recovery())
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	initBAsePlatformRoutes()

	registerCustomValidation()

	for _, fnRegistRoute := range gfnRouter {
		fnRegistRoute(rgPublic, rgAuth)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("start server Error: %s", err.Error()))
			return
		}
	}()

	<-ctx.Done()
	//cancelCtx()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("stop server error: %s", err.Error()))
		return
	}
	global.Logger.Info("stop server success")
}

// 初始化路由
func initBAsePlatformRoutes() {
	InitUserRoutes()
}

// 自定义验证器
func registerCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			value, _ := fl.Field().Interface().(string)
			if value != "" && strings.Index(value, "a") == 0 {
				return true
			}
			return false
		})
	}
}

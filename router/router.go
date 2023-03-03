package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	_ "gokyrie/docs"
	"gokyrie/global"
	"gokyrie/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	r.Use(middleware.Cors())
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	registerCustomValidation()

	initBAsePlatformRoutes()

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
		v.RegisterValidation("mobile", ValidateMobile)
	}
}

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().Interface().(string)
	// 使用正则表达式判断mobile是否合法
	//pattern := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
	pattern := `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`
	ok, _ := regexp.MatchString(pattern, mobile)
	return ok
}

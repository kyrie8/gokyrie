package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//MaxAge:           12 * time.Hour,
	}
	return cors.New(cfg)
}

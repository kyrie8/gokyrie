package middleware

import (
	"gokyrie/api"
	"gokyrie/conf"
	"gokyrie/global"
	"gokyrie/model"
	"gokyrie/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	TOKEN_NAME         = "Authorization"
	TOKEN_PREFIX       = "Bearer "
	TOKEN_ERR_CODE     = 401
	TOKEN_REFRESH_TIME = 10 * 60 * time.Second
)

func TokenErr(ctx *gin.Context) {
	api.Fail(ctx, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   401,
		Msg:    "Invalid Token",
	})
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//no token
		token := ctx.GetHeader(TOKEN_NAME)
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			TokenErr(ctx)
			return
		}

		//token parse err
		token = token[len(TOKEN_PREFIX):]
		iJwtCusterClaims, err := utils.ParseToken(token)
		userId := iJwtCusterClaims.ID
		if err != nil || userId == 0 {
			TokenErr(ctx)
			return
		}

		// token no equal
		stUserId := strconv.Itoa(int(userId))
		redisTokenKey := strings.Replace(conf.LOGIN_USER_REDIS_KEY, "{ID}", stUserId, -1)
		stToken, err := global.RedisClient.Get(redisTokenKey)
		if err != nil || token != stToken {
			TokenErr(ctx)
			return
		}

		// token expiration
		nTokenExpireDuration, err := global.RedisClient.GetExpireDuration(redisTokenKey)
		if err != nil || nTokenExpireDuration <= 0 {
			TokenErr(ctx)
			return
		}

		// refresh token
		if nTokenExpireDuration.Seconds() < TOKEN_REFRESH_TIME.Seconds() {
			nToken, err := GenerateAndCreateUserToken(userId, iJwtCusterClaims.Name)
			if err != nil {
				TokenErr(ctx)
				return
			}
			ctx.Header("token", nToken)
		}

		ctx.Set(conf.LOGIN_USER, model.LoginUser{
			ID:   userId,
			Name: iJwtCusterClaims.Name,
		})
		ctx.Next()
	}
}

func GenerateAndCreateUserToken(userId uint, userName string) (string, error) {
	token, err := utils.GenerateToken(userId, userName)
	if err == nil {
		err = global.RedisClient.Set(strings.Replace(conf.LOGIN_USER_REDIS_KEY, "{ID}", strconv.Itoa(int(userId)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
	}
	return token, err
}

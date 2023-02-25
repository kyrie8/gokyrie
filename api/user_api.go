package api

import (
	"errors"
	"fmt"
	"gokyrie/service/dto"
	"gokyrie/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserApi struct{}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录描述
// @Accept       application/json
// @Produce      application/json
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&iUserLoginDTO)
	if errs != nil {
		Fail(ctx, ResponseJson{
			Msg: parseValidateErrors(errs.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}
	OK(ctx, ResponseJson{
		Msg:  "Login Success",
		Data: iUserLoginDTO,
	})
}

func parseValidateErrors(err validator.ValidationErrors, target any) error {
	var errRes error
	fields := reflect.TypeOf(target).Elem()
	for _, v := range err {
		field, _ := fields.FieldByName(v.Field())
		errMsgTag := fmt.Sprintf("%s_err", v.Tag())
		errMsg := field.Tag.Get(errMsgTag)
		if errMsg == "" {
			errMsg = field.Tag.Get("message")
		}
		if errMsg == "" {
			errMsg = fmt.Sprintf("%s: %s Error", v.Field(), v.Tag())
		}
		errRes = utils.AppendError(errRes, errors.New(errMsg))
	}
	return errRes
}

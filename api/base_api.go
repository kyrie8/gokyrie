package api

import (
	"errors"
	"fmt"
	"gokyrie/global"
	"gokyrie/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx    *gin.Context
	DTO    interface{}
	UriDTO interface{}
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	m.Ctx = option.Ctx

	if option.UriDTO != nil {
		err := m.Ctx.ShouldBindUri(option.UriDTO)
		errResult = utils.AppendError(errResult, err)
	}
	if option.DTO != nil {
		err := m.Ctx.ShouldBind(option.DTO)
		errResult = utils.AppendError(errResult, err)
	}
	if errResult != nil {
		errResult = m.ParseValidateErrors(errResult, option.DTO)
		m.AddError(errResult)
		m.Fail(ResponseJson{
			Msg: m.GetError().Error(),
		})
	}
	return m
}

func (m *BaseApi) AddError(err error) {
	m.Errors = utils.AppendError(m.Errors, err)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}

// 全部错误返回
// func (m *BaseApi) ParseValidateErrors(err error, target any) error {
// 	var errRes error
// 	fields := reflect.TypeOf(target).Elem()
// 	fmt.Printf("fields: %v\n", fields)
// 	validatorErr, ok := err.(validator.ValidationErrors)
// 	fmt.Printf("validatorErr: %v\n", validatorErr)
// 	if !ok {
// 		return err
// 	}
// 	for _, v := range validatorErr {
// 		fmt.Printf("v: %v\n", v)
// 		field, _ := fields.FieldByName(v.Field())
// 		errMsgTag := fmt.Sprintf("%s_err", v.Tag())
// 		errMsg := field.Tag.Get(errMsgTag)
// 		if errMsg == "" {
// 			errMsg = field.Tag.Get("message")
// 		}
// 		if errMsg == "" {
// 			errMsg = fmt.Sprintf("%s: %s Error", v.Field(), v.Tag())
// 		}
// 		errRes = utils.AppendError(errRes, errors.New(errMsg))
// 	}
// 	return errRes
// }

func (m *BaseApi) ParseValidateErrors(err error, target any) error {
	var errRes error
	fields := reflect.TypeOf(target).Elem()
	validatorErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	for _, fieldErr := range validatorErr {
		field, _ := fields.FieldByName(fieldErr.Field())
		errTag := fieldErr.Tag() + "_err"
		fmt.Printf("%v", fieldErr.Tag())
		errMsg := field.Tag.Get(errTag)
		if errMsg == "" {
			errMsg = field.Tag.Get("message")
		}
		if errMsg != "" {
			errRes = errors.New(errMsg)
			break
		}
	}
	return errRes
}

func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}

func (m *BaseApi) OK(resp ResponseJson) {
	OK(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(m.Ctx, resp)
}

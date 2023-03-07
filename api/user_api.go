package api

import (
	"fmt"
	"gokyrie/conf"
	"gokyrie/service"
	"gokyrie/service/dto"
	"gokyrie/utils"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录描述
// @Accept       application/json
// @Produce      application/json
// @Param body body dto.UserLoginDTO true "body"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.Fail(ResponseJson{
			Msg:  err.Error(),
			Code: conf.FAIL_CODE,
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	u.OK(ResponseJson{
		Msg:  "Login Success",
		Code: conf.SUCCESS_CODE,
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

// @Tag 用户管理
// @Summary 用户添加
// @Description 用户添加描述
// @Accept       application/json
// @Produce      application/json
// @Param body body dto.UserAddDTO true "body"
// @Router /api/v1/user [post]
func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}
	file, err := c.FormFile("file")
	if err == nil {
		stFilePath := fmt.Sprintf("./upload/%s", file.Filename)
		_ = c.SaveUploadedFile(file, stFilePath)
		iUserAddDTO.Avatar = stFilePath
	}
	err = m.Service.AddUser(&iUserAddDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iUserAddDTO,
		Code: conf.SUCCESS_CODE,
	})
}

func (m UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	iUser, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iUser,
		Code: conf.SUCCESS_CODE,
	})
}

func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDto dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDto}).GetError(); err != nil {
		return
	}
	giUserList, noTotal, err := m.Service.GetUserList(&iUserListDto)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  giUserList,
		Total: noTotal,
		Code:  conf.SUCCESS_CODE,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	err := m.Service.UpdateUser(&iUserUpdateDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Code: conf.SUCCESS_CODE,
	})
}

func (m UserApi) DeleteUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	err := m.Service.DeleteUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Code: conf.SUCCESS_CODE,
	})
}

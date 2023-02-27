package api

import (
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
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	u.OK(ResponseJson{
		Msg: "Login Success",
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

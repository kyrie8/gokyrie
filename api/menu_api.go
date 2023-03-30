package api

import (
	"gokyrie/conf"
	"gokyrie/service"
	"gokyrie/service/dto"

	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	BaseApi
	Service *service.MenuService
}

func NewMenuApi() MenuApi {
	return MenuApi{
		BaseApi: NewBaseApi(),
		Service: service.NewMenuService(),
	}
}

func (m MenuApi) AddMenu(c *gin.Context) {
	var iMenuDto dto.MenuAddDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMenuDto}).GetError(); err != nil {
		return
	}
	err := m.Service.AddMenu(&iMenuDto)
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

func (m MenuApi) UpdateMenu(c *gin.Context) {
	var iMenuUpdateDTO dto.MenuUpdateDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMenuUpdateDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	err := m.Service.UpdateMenu(&iMenuUpdateDTO)
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

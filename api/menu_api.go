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
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMenuUpdateDTO, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	iMenuUpdateDTO.MenuId = iCommonIDDTO.ID
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

func (m MenuApi) GetMenuList(c *gin.Context) {
	var iMenuListDTO dto.MenuListDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iMenuListDTO}).GetError(); err != nil {
		return
	}
	iMenu, total, err := m.Service.GetMenuList(&iMenuListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	iMenu = m.Service.GetMenuTree(iMenu, 0)
	m.OK(ResponseJson{
		Code:  conf.SUCCESS_CODE,
		Data:  iMenu,
		Total: total,
	})
}

package api

import (
	"gokyrie/conf"
	"gokyrie/service"
	"gokyrie/service/dto"

	"github.com/gin-gonic/gin"
)

type RoleApi struct {
	BaseApi
	Service *service.RoleService
}

func NewRoleApi() RoleApi {
	return RoleApi{
		BaseApi: NewBaseApi(),
		Service: service.NewRoleService(),
	}
}

func (m RoleApi) AddRole(c *gin.Context) {
	var iRoleDto dto.RoleAddDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iRoleDto}).GetError(); err != nil {
		return
	}
	err := m.Service.AddRole(&iRoleDto)
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

func (m RoleApi) UpdateRole(c *gin.Context) {
	var iRoleUpdateDTO dto.RoleUpdateDto
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iRoleUpdateDTO, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	iRoleUpdateDTO.RoleId = iCommonIDDTO.ID
	err := m.Service.UpdateRole(&iRoleUpdateDTO)
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

func (m RoleApi) GetRoleList(c *gin.Context) {
	var iRoleListDto dto.RoleListDto
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iRoleListDto}).GetError(); err != nil {
		return
	}
	iRole, total, err := m.Service.GetRoleList(&iRoleListDto)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Code:  conf.SUCCESS_CODE,
		Data:  iRole,
		Total: total,
	})
}

func (m RoleApi) UpdateRoleMenu(c *gin.Context) {
	var iRoleMenuDTO dto.RoleMenuDto
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iRoleMenuDTO, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	iRoleMenuDTO.RoleId = iCommonIDDTO.ID
	err := m.Service.RoleMenu(&iRoleMenuDTO)
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

func (m RoleApi) GetMenuByRoleId(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	var menuId []uint
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	iRole, err := m.Service.GetMenusByRoleId(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	for _, v := range iRole.Menus {
		menuId = append(menuId, v.MenuId)
	}
	m.OK(ResponseJson{
		Code: conf.SUCCESS_CODE,
		Data: menuId,
	})
}

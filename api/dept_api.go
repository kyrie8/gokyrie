package api

import (
	"gokyrie/conf"
	"gokyrie/service"
	"gokyrie/service/dto"

	"github.com/gin-gonic/gin"
)

type DeptApi struct {
	BaseApi
	Service *service.DeptService
}

func NewDeptApi() DeptApi {
	return DeptApi{
		BaseApi: NewBaseApi(),
		Service: service.NewDeptService(),
	}
}

func (m DeptApi) DeleteDeptById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	err := m.Service.DeleteDeptById(&iCommonIDDTO)
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

func (m DeptApi) AddDept(c *gin.Context) {
	var iDeptDto dto.DeptAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iDeptDto}).GetError(); err != nil {
		return
	}
	err := m.Service.AddDept(&iDeptDto)
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

func (m DeptApi) UpdateDept(c *gin.Context) {
	var iDeptUpdateDto dto.DeptUpdateDTO
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iDeptUpdateDto, UriDTO: &iCommonIDDTO}).GetError(); err != nil {
		return
	}
	iDeptUpdateDto.DeptId = iCommonIDDTO.ID
	err := m.Service.UpdateDept(&iDeptUpdateDto)
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

func (m DeptApi) GetDeptList(c *gin.Context) {
	var iDeptListDto dto.DeptListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iDeptListDto}).GetError(); err != nil {
		return
	}
	iDept, total, err := m.Service.GetDeptList(&iDeptListDto)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: conf.FAIL_CODE,
			Msg:  err.Error(),
		})
		return
	}
	iDept = m.Service.GetDeptTree(iDept, 0)
	m.OK(ResponseJson{
		Code:  conf.SUCCESS_CODE,
		Data:  iDept,
		Total: total,
	})
}

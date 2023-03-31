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
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: iCommonIDDTO}).GetError(); err != nil {
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

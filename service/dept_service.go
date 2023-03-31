package service

import (
	"errors"
	"gokyrie/dao"
	"gokyrie/model"
	"gokyrie/service/dto"
)

type DeptService struct {
	BaseService
	Dao *dao.DeptDao
}

var deptService *DeptService

func NewDeptService() *DeptService {
	if deptService == nil {
		deptService = &DeptService{
			Dao: dao.NewDeptDao(),
		}
	}
	return deptService
}

func (m *DeptService) AddDept(iDeptDto *dto.DeptAddDTO) error {
	if m.Dao.CheckDeptNameExist(iDeptDto.DeptName) {
		return errors.New("DeptName exist")
	}
	return m.Dao.AddDept(iDeptDto)
}

func (m *DeptService) UpdateDept(iDeptDto *dto.DeptUpdateDTO) error {
	if iDeptDto.DeptId <= 0 {
		return errors.New("invalid deptId")
	}
	return m.Dao.UpdateDept(iDeptDto)
}

func (m *DeptService) GetDeptList(iDeptListDto *dto.DeptUserListDTO) ([]model.Dept, int64, error) {
	return m.Dao.GetDeptList(iDeptListDto)
}

func (m *DeptService) DeleteDeptById(iCommonIDDTO *dto.CommonIDDTO) error {
	return m.Dao.DeleteDeptById(iCommonIDDTO.ID)
}

package service

import (
	"errors"
	"gokyrie/dao"
	"gokyrie/model"
	"gokyrie/service/dto"
)

type RoleService struct {
	BaseService
	Dao *dao.RoleDao
}

var roleService *RoleService

func NewRoleService() *RoleService {
	if roleService == nil {
		roleService = &RoleService{
			Dao: dao.NewRoleDao(),
		}
	}
	return roleService
}

func (m *RoleService) AddRole(iRoleDto *dto.RoleAddDto) error {
	if m.Dao.CheckRoleNameExist(iRoleDto.RoleName) {
		return errors.New("roleName exist")
	}
	return m.Dao.AddRole(iRoleDto)
}

func (m *RoleService) UpdateRole(iRoleDto *dto.RoleUpdateDto) error {
	if iRoleDto.RoleId <= 0 {
		return errors.New("invalid menuId")
	}
	return m.Dao.UpdateRole(iRoleDto)
}

func (m *RoleService) GetRoleList(iRoleListDto *dto.RoleListDto) ([]model.Role, int64, error) {
	return m.Dao.GetRoleList(iRoleListDto)
}

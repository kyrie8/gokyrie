package dao

import (
	"gokyrie/model"
	"gokyrie/service/dto"
)

type RoleDao struct {
	BaseDao
}

var roleDao *RoleDao

func NewRoleDao() *RoleDao {
	if roleDao == nil {
		roleDao = &RoleDao{
			BaseDao: NewBaseDao(),
		}
	}
	return roleDao
}

func (m *RoleDao) GetUserByName(name string) (model.Role, error) {
	var iRole model.Role
	err := m.Orm.Model(&model.Role{}).Where("name=?", name).Find(&iRole).Error
	return iRole, err
}

func (m *RoleDao) CheckRoleNameExist(name string) bool {
	var nTotal int64
	m.Orm.Model(&model.Role{}).Where("name = ?", name).Count(&nTotal)
	return nTotal > 0
}

func (m *RoleDao) AddRole(iRoleDto *dto.RoleAddDto) error {
	var iRole model.Role
	iRoleDto.ConvertToModel(&iRole)
	result := m.Orm.Create(&iRole)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *RoleDao) UpdateRole(iRoleUpdateDto *dto.RoleUpdateDto) error {
	var iRole *model.Role
	return m.Orm.Model(iRole).Where("menu_id = ?", iRoleUpdateDto.RoleId).Select("*").Omit("CreatedAt").Updates(iRoleUpdateDto).Error
}

func (m *RoleDao) GetRoleList(dto *dto.RoleListDto) ([]model.Role, int64, error) {
	var nTotal int64
	var iRole []model.Role
	Db := m.Orm
	if dto.Name != "" {
		Db = Db.Where("role_name = ?", dto.Name)
	}
	err := Db.Scopes(Paginate(dto.Paginate)).Find(&iRole).Offset(-1).Limit(-1).Count(&nTotal).Error
	return iRole, nTotal, err
}

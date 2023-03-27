package dao

import (
	"gokyrie/model"
	"gokyrie/service/dto"
)

type MenuDao struct {
	BaseDao
}

var menuDao *MenuDao

func NewMenuDao() *MenuDao {
	if menuDao == nil {
		menuDao = &MenuDao{
			BaseDao: NewBaseDao(),
		}
	}
	return menuDao
}

func (m *MenuDao) AddMenu(iMenuDto *dto.MenuAddDto) error {
	var iMenu model.Menu
	iMenuDto.ConvertToModel(&iMenu)
	result := m.Orm.Create(&iMenu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MenuDao) CheckMenuNameExist(name string) bool {
	var noTotal int64
	m.Orm.Model(&model.Menu{}).Where("menuName = ?", name).Count(&noTotal)
	return noTotal > 0
}

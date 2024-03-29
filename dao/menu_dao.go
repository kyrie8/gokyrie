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
	m.Orm.Model(&model.Menu{}).Where("menu_name = ?", name).Count(&noTotal)
	return noTotal > 0
}

func (m *MenuDao) UpdateMenu(iMenuUpdateDto *dto.MenuUpdateDto) error {
	var iMenu *model.Menu
	return m.Orm.Model(iMenu).Where("menu_id = ?", iMenuUpdateDto.MenuId).Select("*").Omit("CreatedAt").Updates(iMenuUpdateDto).Error
}

func (m *MenuDao) GetMenuList(dto *dto.MenuListDto) ([]model.Menu, int64, error) {
	var nTotal int64
	var iMenu []model.Menu
	Db := m.Orm
	if dto.Name != "" {
		Db = Db.Where("menu_name = ?", dto.Name)
	}
	err := Db.Scopes(Paginate(dto.Paginate)).Find(&iMenu).Offset(-1).Limit(-1).Count(&nTotal).Error
	return iMenu, nTotal, err
}

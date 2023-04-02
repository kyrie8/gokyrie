package service

import (
	"errors"
	"gokyrie/dao"
	"gokyrie/model"
	"gokyrie/service/dto"
)

type MenuService struct {
	BaseService
	Dao *dao.MenuDao
}

var menuService *MenuService

func NewMenuService() *MenuService {
	if menuService == nil {
		menuService = &MenuService{
			Dao: dao.NewMenuDao(),
		}
	}
	return menuService
}

func (m *MenuService) AddMenu(iMenuDto *dto.MenuAddDto) error {
	if m.Dao.CheckMenuNameExist(iMenuDto.MenuName) {
		return errors.New("menuName exist")
	}
	return m.Dao.AddMenu(iMenuDto)
}

func (m *MenuService) UpdateMenu(iMenuDto *dto.MenuUpdateDto) error {
	if iMenuDto.MenuId <= 0 {
		return errors.New("invalid menuId")
	}
	return m.Dao.UpdateMenu(iMenuDto)
}

func (m *MenuService) GetMenuList(iMenuListDto *dto.MenuListDto) ([]model.Menu, int64, error) {
	return m.Dao.GetMenuList(iMenuListDto)
}

func (m *MenuService) GetMenuTree(array []model.Menu, pid uint) []model.Menu {
	var res []model.Menu
	for _, v := range array {
		if v.ParentId == pid {
			v.Children = m.GetMenuTree(array, v.MenuId)
			res = append(res, v)
		}
	}
	return res
}

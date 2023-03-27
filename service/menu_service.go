package service

import (
	"errors"
	"gokyrie/dao"
	"gokyrie/service/dto"
)

type MenuService struct {
	BaseService,
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
		return errors.New("MenuName exist")
	}
	return m.Dao.AddMenu(iMenuDto)
}

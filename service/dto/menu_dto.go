package dto

import "gokyrie/model"

type MenuAddDto struct {
	ParentId  uint   `json:"parentId" form:"parentId"`
	MenuName  string `json:"menuName" form:"menuName" binding:"required" message:"菜单名不能为空"`
	Path      string `json:"path" form:"path" binding:"required" message:"路径不能为空"`
	Component string `json:"component" form:"component" binding:"required" message:"组件不能为空"`
	Icon      string `json:"icon" form:"icon"`
	Type      uint   `json:"type" form:"type"`
	IsOutLink uint   `json:"isOutLink" form:"isOutLink"`
	Hidden    uint   `json:"hidden" form:"hidden"`
	AuthKey   string `json:"authKey" form:"authKey"`
}

func (m *MenuAddDto) ConvertToModel(iMenu *model.Menu) {
	iMenu.AuthKey = m.AuthKey
	iMenu.Hidden = &m.Hidden
	iMenu.IsOutLink = &m.IsOutLink
	iMenu.Type = &m.Type
	iMenu.Icon = m.Icon
	iMenu.Component = m.Component
	iMenu.Path = m.Path
	iMenu.MenuName = m.MenuName
	iMenu.ParentId = &m.ParentId
}

type MenuUpdateDto struct {
	ParentId  uint   `json:"parentId" form:"parentId"`
	MenuName  string `json:"menuName" form:"menuName" binding:"required" message:"菜单名不能为空"`
	Path      string `json:"path" form:"path" binding:"required" message:"路径不能为空"`
	Component string `json:"component" form:"component" binding:"required" message:"组件不能为空"`
	Icon      string `json:"icon" form:"icon"`
	Type      uint   `json:"type" form:"type"`
	IsOutLink uint   `json:"isOutLink" form:"isOutLink"`
	Hidden    uint   `json:"hidden" form:"hidden"`
	AuthKey   string `json:"authKey" form:"authKey"`
	MenuId    uint   `json:"menuId" form:"menuId" uri:"menuId"`
}

func (m *MenuUpdateDto) ConvertToModel(iMenu *model.Menu) {
	iMenu.AuthKey = m.AuthKey
	iMenu.Hidden = &m.Hidden
	iMenu.IsOutLink = &m.IsOutLink
	iMenu.Type = &m.Type
	iMenu.Icon = m.Icon
	iMenu.Component = m.Component
	iMenu.Path = m.Path
	iMenu.MenuName = m.MenuName
	iMenu.ParentId = &m.ParentId
	iMenu.MenuId = m.MenuId
}

package dto

import "gokyrie/model"

type RoleAddDto struct {
	RoleName string `json:"roleName" form:"roleName"`
	Remark   string `json:"remark" form:"remark"`
}

func (m *RoleAddDto) ConvertToModel(iRole *model.Role) {
	iRole.RoleName = m.RoleName
	iRole.Remark = m.Remark
}

type RoleListDto struct {
	Paginate
	Name string `json:"name" form:"name"`
}

type RoleUpdateDto struct {
	RoleAddDto
	RoleId uint `json:"roleId" form:"roleId"`
}

package dto

import (
	"gokyrie/model"
	"gokyrie/utils"
)

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"realName" form:"realName"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile" binding:"omitempty,mobile" message:"手机号不正确"`
	Email    string `json:"email" form:"email" binding:"omitempty,email" message:"邮箱不正确"`
	DeptId   uint   `json:"dept_id" form:"dept_id" binding:"omitempty"`
	RoleId   []uint `json:"role_id" form:"role_id" binding:"omitempty"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	stHash, _ := utils.Encrypt(m.Password)
	iUser.Email = m.Email
	iUser.Password = stHash
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.Avatar = m.Avatar
	if m.DeptId > 0 {
		iUser.DeptId = m.DeptId
	}
}

type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id" uri:"id"`
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"realName" form:"realName"`
	Mobile   string `json:"mobile" form:"mobile" binding:"omitempty,mobile" message:"手机号不正确"`
	Email    string `json:"email" form:"email" binding:"omitempty,email" message:"邮箱不正确"`
	RoleId   []uint `json:"role_id" form:"role_id"` //暂时允许角色为空
	DeptId   uint   `json:"dept_id" form:"dept_id"`
}

func (m *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.Email = m.Email
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.ID = m.ID
	if m.DeptId > 0 {
		iUser.DeptId = m.DeptId
	}
}

type UserListDTO struct {
	Paginate
	Name string `json:"name" form:"name"`
}

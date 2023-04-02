package dto

import "gokyrie/model"

type DeptAddDTO struct {
	ParentId uint   `json:"parentId" form:"parentId"`
	DeptName string `json:"deptName" form:"deptName" binding:"required"`
	Remark   string `json:"remark" form:"remark"`
}

func (m *DeptAddDTO) ConvertToModel(iDept *model.Dept) {
	iDept.ParentId = m.ParentId
	iDept.DeptName = m.DeptName
	iDept.Remark = m.Remark
}

type DeptUpdateDTO struct {
	DeptAddDTO
	DeptId uint `json:"deptId" form:"deptId" uri:"deptId"`
}

type DeptListDTO struct {
	Paginate
	Name string `json:"name" form:"name"`
}

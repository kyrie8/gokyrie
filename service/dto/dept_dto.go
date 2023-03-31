package dto

type DeptAddDTO struct {
	ParentId uint   `json:"parentId" form:"parentId"`
	DeptName string `json:"deptName" form:"deptName" binding:"required"`
	Remark   string `json:"remark" form:"remark"`
}

type DeptUpdateDTO struct {
	DeptAddDTO
	DeptId uint `json:"deptId" form:"deptId" uri:"deptId"`
}

type DeptUserListDTO struct {
	Paginate
	Name string `json:"name" form:"name"`
}

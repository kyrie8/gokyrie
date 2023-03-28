package model

type Dept struct {
	DeptId   uint   `json:"deptId" gorm:"primaryKey;autoIncrement"`
	ParentId uint   `json:"parentId" gorm:"default:0"`
	DeptName string `json:"deptName" gorm:"size:255;not null"`
	Remark   string `json:"remark" gorm:"size:255;comment:'备注说明'"`
	Users    []User `json:"users,omitempty" gorm:"foreignKey:DeptId"`
	Common
}

func (Dept) tableName() string {
	return "t_dept"
}

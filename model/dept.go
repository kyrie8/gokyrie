package model

type Dept struct {
	DeptId   uint   `json:"dept_id" gorm:"primaryKey;autoIncrement"`
	ParentId uint   `json:"-" gorm:"default:0"`
	DeptName string `json:"dept_name" gorm:"size:255;not null"`
	Remark   string `json:"remark" gorm:"size:255;comment:'备注说明'"`
	Users    []User `gorm:"foreignKey:DeptId"`
	Common
}
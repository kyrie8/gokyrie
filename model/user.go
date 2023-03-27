package model

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"realName" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Mobile   string `json:"mobile" gorm:"size:11"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"size:128;not null"`
	Roles    []Role `gorm:"many2many:user_role;"`
	DeptId   uint   `json:"deptId"`
	gorm.Model
}

// func (m *User) Encrypt() error {
// 	stHash, err := utils.Encrypt(m.Password)
// 	if err == nil {
// 		m.Password = stHash
// 	}
// 	return err
// }

// func (m *User) BeforeCreate(orm *gorm.DB) error {
// 	return m.Encrypt()
// }

type LoginUser struct {
	ID   uint
	Name string
}

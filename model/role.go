package model

type Role struct {
	RoleId   uint   `json:"roleId" gorm:"primaryKey;autoIncrement"`
	RoleName string `json:"roleName" gorm:"size:255;not null"`
	Remark   string `json:"remark" gorm:"size:255;comment:'角色备注'"`
	Menus    []Menu `json:"menus" gorm:"many2many:role_menu;foreignKey:RoleId;joinForeignKey:RoleId;references:MenuId;joinReferences:MenuId"`
	Users    []User `json:"users" gorm:"many2many:user_role;foreignKey:RoleId;references:id;joinForeignKey:RoleId;"`
	Common
}

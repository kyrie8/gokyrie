package model

type Role struct {
	RoleId   uint   `json:"role_id" gorm:"primaryKey;autoIncrement"`
	RoleName string `json:"role_name" gorm:"size:255;not null"`
	Remark   string `json:"remark" gorm:"size:255;comment:'角色备注'"`
	Menus    []Menu `gorm:"many2many:role_menu;foreignKey:RoleId;joinForeignKey:RoleId;references:MenuId;joinReferences:MenuId"`
	Users    []User `gorm:"many2many:user_role;foreignKey:RoleId;references:id;joinForeignKey:RoleId;"`
	Common
}

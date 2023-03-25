package model

type Menu struct {
	MenuId    uint   `json:"menu_id" gorm:"primaryKey;autoIncrement"`
	ParentId  uint   `json:"-" gorm:"default:0"`
	MenuName  string `json:"menu_name" gorm:"size:20;not null"`
	Path      string `json:"path" gorm:"size:255;not null;comment:'路由地址'"`
	Component string `json:"component" gorm:"size:255;not null;comment:'组件路径'"`
	Icon      string `json:"icon" gorm:"size:20"`
	Type      uint   `json:"type" gorm:"default:0;comment:'0是菜单,1是按钮'"`
	IsOutLink uint   `json:"is_out_link" gorm:"default:0;comment:'0不是外链,1是外链'"`
	Hidden    uint   `json:"hidden" gorm:"default:0;comment:'0不隐藏,1隐藏'"`
	AuthKey   string `json:"auth_key" gorm:"not null;comment:'权限标识:如user:add'"`
	Roles     []Role `gorm:"many2many:role_menu;foreignKey:MenuId;joinForeignKey:MenuId;references:RoleId;joinReferences:RoleId"`
	Common
}

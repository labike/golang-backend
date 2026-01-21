package entity

import "go-admin/util"

type SysRole struct {
	ID          uint       `gorm:"column:id;comment:'id';NOT NULL" json:"id"`
	RoleName    string     `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`
	RoleKey     string     `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`
	Status      int        `gorm:"column:status;default:1;comment:'账号状态1->启用 2->禁用';NOT NULL" json:"status"`
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

type AddSysRoleDto struct {
	RoleName    string
	RoleKey     string
	Status      int
	Description string
}

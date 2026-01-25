package entity

// 用户与角色关系模型
type SysAdminRole struct {
	AdminId int `gorm:"column:admin_id;comment:'用户id';NOT NULL" json:"adminId"`
	RoleId  int `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}

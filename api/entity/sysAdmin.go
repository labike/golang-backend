package entity

import "go-admin/util"

// 用户模型对象
type SysAdmin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        //ID
	PostId     int        `gorm:"column:post_id;comment:'岗位id'" json:"postId"`                                 // 岗位id
	DeptId     int        `gorm:"column:dept_id;comment:'部门id'" json:"deptId"`                                 // 部门id
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`           // 密码
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`                    // 昵称
	Status     int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`                           //  头像
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`                          // 邮箱
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`                          // 电话
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`                           // 备注
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结构体
type JwtAdmin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"` // 备注
}

// 登录对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image" validate:"required,min=4,max=6"` // 验证码
	IdKey    string `json:"idKey" validate:"required"`             // uuid
}

type AddSysAdminDto struct {
	PostId   int    `validate:"required"`
	DeptId   int    `validate:"required"`
	RoleId   int    `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string
	Status   int `validate:"required"`
}

type SysAdminInfo struct {
	ID       uint   `json:"id"`
	PostId   int    `json:"postId"`
	DeptId   int    `json:"deptId"`
	RoleId   int    `json:"roleId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Note     string `json:"note"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

type UpdateSysAdminDto struct {
	Id       uint
	PostId   int
	DeptId   int
	RoleId   int
	Username string
	Nickname string
	Phone    string
	Email    string
	Note     string
	Status   int
}

type SysAdminIdDto struct {
	Id uint `json:"id"`
}

type UpdateSysAdminStatusDto struct {
	Status int
	Id     uint
}

type ResetSysAdminPasswordDto struct {
	Id       uint
	Password string
}

type SysAdminVo struct {
	ID         uint       `json:"id"`
	PostId     int        `json:"postId"`
	DeptId     int        `json:"deptId"`
	RoleId     int        `json:"roleId"`
	Username   string     `json:"username"`
	Nickname   string     `json:"nickname"`
	PostName   string     `json:"postName"`
	RoleName   string     `json:"roleName"`
	DeptName   string     `json:"deptName"`
	Icon       string     `json:"icon"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Note       string     `json:"note"`
	Status     int        `json:"status"`
	CreateTime util.HTime `json:"createTime"`
}

type UpdateUserDto struct {
	Id       uint
	Icon     string
	Username string
	Nickname string
	Phone    string
	Email    string
	Note     string
}

type UpdateUserPasswordDto struct {
	Id            uint
	Password      string `validate:"required"`
	NewPassword   string `validate:"required"`
	ResetPassword string `validate:"required"`
}

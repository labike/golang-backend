package entity

import "go-admin/util"

type SysOperateLog struct {
	ID         uint       `gorm:"column:id;comment:'id';primaryKey;NOT NULL" json:"id"`
	AdminId    uint       `gorm:"column:admin_id;comment:'管理员id';NOT NULL" json:"adminId"`
	Username   string     `gorm:"column:username;varchar(64);comment:'管理员账号';" json:"username"`
	Method     string     `gorm:"column:method;varchar(64);comment:'请求方式';" json:"method"`
	Ip         string     `gorm:"column:ip;varchar(64);comment:'ip';" json:"ip"`
	Url        string     `gorm:"column:url;varchar(500);comment:'Url';" json:"url"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysOperateLog) TableName() string {
	return "sys_operation_log"
}

type SysOperateIdDto struct {
	Id uint `json:"id"`
}

type BatchSysOperateDto struct {
	Ids []uint
}

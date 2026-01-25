package entity

import "go-admin/util"

type SysLoginLog struct {
	ID            uint       `gorm:"column:id;comment:'id';primaryKey;NOT NULL" json:"id"`
	Username      string     `gorm:"column:username;varchar(50);comment:'用户账号';" json:"username"`
	IpAddress     string     `gorm:"column:ip_address;varchar(128);comment:'登录ip地址';" json:"ipAddress"`
	LoginLocation string     `gorm:"column:login_location;varchar(255);comment:'登陆地址';" json:"loginLocation"`
	Browser       string     `gorm:"column:browser;varchar(50);comment:'登录浏览器';" json:"browser"`
	Os            string     `gorm:"column:os;varchar(50);comment:'操作系统';" json:"os"`
	LoginStatus   int        `gorm:"column:login_status;comment:'登陆状态(1->成功 2->失败)';" json:"loginStatus"`
	Message       string     `gorm:"column:message;varchar(255);comment:'提示消息';" json:"message"`
	LoginTime     util.HTime `gorm:"column:login_time;comment:'登录时间';" json:"loginTime"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_info"
}

type SysLoginInfoIdDto struct {
	Id uint `json:"id"`
}

type DeleteLoginInfoDto struct {
	Ids []uint
}

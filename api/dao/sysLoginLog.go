package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
	"go-admin/util"
	"time"
)

func CreateLoginlog(
	username,
	ipAddress,
	loginLocation,
	browser,
	os,
	message string,
	loginStatus int) {
	sysLoginLog := entity.SysLoginLog{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     util.HTime{Time: time.Now()},
	}
	Db.Save(&sysLoginLog)
}

func GetSysLoginLogList(
	Username,
	LoginStatus,
	BeginTime,
	EndTime string,
	PageNum,
	PageSize int) (sysLoginInfo []entity.SysLoginLog, count int64) {
	curDb := Db.Table("sys_login_info")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}
	if LoginStatus != "" {
		curDb = curDb.Where("login_status = ?", LoginStatus)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("login_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("login_time DESC").Find(&sysLoginInfo)
	return sysLoginInfo, count
}

func BatchDeleteSysLoginLog(dto entity.DeleteLoginInfoDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysLoginLog{})
}

func DeleteSysLoginLog(dto entity.SysLoginInfoIdDto) {
	Db.Delete(&entity.SysLoginLog{}, dto.Id)
}

func ClearSysLoginLog() {
	Db.Exec("truncate table sys_login_info")
}

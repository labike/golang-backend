package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
)

func CreateSysOperateLog(log entity.SysOperateLog) {
	Db.Create(&log)
}

func GetSysOperateList(
	Username,
	BeginTime,
	EndTime string,
	PageSize,
	PageNum int) (operaList []entity.SysOperateLog, count int64) {
	curDb := Db.Table("sys_operation_log")
	if Username != "" {
		curDb = curDb.Where("Username = ?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("BeginTime BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&operaList)
	return operaList, count
}

func DeleteOperateLog(dto entity.SysOperateIdDto) {
	Db.Delete(&entity.SysOperateLog{}, dto)
}

func BatchDeleteOperateLog(dto entity.BatchSysOperateDto) {
	Db.Where("id in (?)", dto.Ids).Delete(entity.SysOperateLog{})
}

func CleanSysOperateLog() {
	Db.Exec("truncate table sys_operation_log")
}

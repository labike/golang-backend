package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
)

// 用户数据层
func SysAdmin(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
	"go-admin/util"
	"time"
)

func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

func GetSysRoleKey(roleKey string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

func CreateSysRole(dto entity.AddSysRoleDto) bool {
	sysRoleName := GetSysRoleByName(dto.RoleName)
	if sysRoleName.ID > 0 {
		return false
	}
	sysRoleKey := GetSysRoleKey(dto.RoleKey)
	if sysRoleKey.ID > 0 {
		return false
	}
	addSysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Status:      dto.Status,
		Description: dto.Description,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

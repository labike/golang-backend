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

func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	Db.First(&sysRole, Id)
	return sysRole
}

func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	return sysRole
}

func DeleteSysRoleById(dto entity.SysRoleIdDto) {
	Db.Table("sys_role").Delete(&entity.SysRole{}, dto.Id)
	Db.Table("sys_role_menu").Where("role_id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
}

func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatusDto) bool {
	var sysRole entity.SysRole
	Db.First(&sysRole, dto.Id)
	sysRole.Status = dto.Status
	tx := Db.Save(&sysRole)
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

// 分页查询
func GetSysRoleList(
	PageNum,
	PageSize int,
	RoleName,
	Status,
	BeginTime,
	EndTime string) (sysRole []entity.SysRole, count int64) {
	curDb := Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	return sysRole, count
}

func GetSysRoleVo() (sysRoleVo []entity.SysRoleVo) {
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

// 根据角色id查询菜单权限
func QueryRoleMenuIdList(Id int) (idVo []entity.IdVo) {
	const menuType int = 3
	// select sm.id from sys_menu sm left join sys_menu srm on srm.menu_id = sm.id
	// left join sys_role sr on sr.id = srm.role_id where sm.menu_type = 3 and
	// sr.id = 1
	Db.Table("sys_menu sm").
		Select("sm.id").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)
	return idVo
}

// 分配权限
func AssignPermission(menu entity.RoleMenu) (err error) {
	err = Db.Table("sys_role_menu").Where("role_id = ?", menu.Id).Delete(&entity.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	for _, value := range menu.MenuIds {
		var entity entity.SysRoleMenu
		entity.RoleId = menu.Id
		entity.MenuId = value
		Db.Create(&entity)
	}
	return err
}

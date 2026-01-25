package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
	"go-admin/util"
	"time"
)

// 用户数据层
func SysAdmin(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

func GetSysAdminByUsername(username string) (sysAdmin entity.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	sysAdminByUserName := GetSysAdminByUsername(dto.Username)
	if sysAdminByUserName.ID > 0 {
		return false
	}
	sysAdmin := entity.SysAdmin{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		Nickname:   dto.Nickname,
		DeptId:     dto.DeptId,
		PostId:     dto.PostId,
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&sysAdmin)
	sysAdminExists := GetSysAdminByUsername(dto.Username)
	var entity entity.SysAdminRole
	entity.AdminId = int(sysAdminExists.ID)
	entity.RoleId = dto.RoleId
	Db.Create(&entity)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetAdminInfoById(id int) (sysAdmin entity.SysAdminInfo) {
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdmin, id)
	return sysAdmin
}

func UpdateSysAdmin(dto entity.UpdateSysAdminDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	sysAdmin.PostId = dto.PostId
	sysAdmin.DeptId = dto.DeptId
	sysAdmin.Status = dto.Status
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	Db.Save(&sysAdmin)
	var sysAdminRole entity.SysAdminRole
	// 先删除之前的角色在分配新的
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
	sysAdminRole.AdminId = int(dto.Id)
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)
	return sysAdmin
}

func DeleteSysAdmin(dto entity.SysAdminIdDto) {
	Db.First(&entity.SysAdmin{}, dto.Id)
	Db.Delete(&entity.SysAdmin{}, dto.Id)
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
}

func UpdateSysAdminStatus(dto entity.UpdateSysAdminStatusDto) {
	var sysAdmin entity.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Status = dto.Status
	Db.Save(&sysAdmin)
}

func ResetSysAdminPassword(dto entity.ResetSysAdminPasswordDto) {
	var sysAdmin entity.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = util.EncryptionMd5(dto.Password)
	Db.Save(&sysAdmin)
}

func GetAdminList(
	PageNum,
	PageSize int,
	Username,
	Status,
	BeginTime,
	EndTime string) (sysAdminVo []entity.SysAdminVo, count int64) {
	curDb := Db.Table("sys_admin").
		Select("sys_admin.*, sys_post.post_name, sys_role.role_name, sys_dept.dept_name").
		Joins("LEFT JOIN sys_post ON sys_admin.post_id = sys_post.id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_dept ON sys_dept.id = sys_admin.dept_id")
	if Username != "" {
		curDb = curDb.Where("sys_admin.username = ?", Username)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	return sysAdminVo, count
}

func UpdateUser(dto entity.UpdateUserDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Icon != "" {
		sysAdmin.Icon = dto.Icon
	}
	Db.Save(&sysAdmin)
	return sysAdmin
}

func UpdateUserPassword(dto entity.UpdateUserPasswordDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = dto.NewPassword
	Db.Save(&sysAdmin)
	return sysAdmin
}

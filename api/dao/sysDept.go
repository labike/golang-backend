package dao

import (
	"go-admin/api/entity"
	"go-admin/util"
	"time"
)
import . "go-admin/pkg/db"

func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []entity.SysDept) {
	curDb := Db.Table("sys_dept")
	if DeptName != "" {
		curDb = Db.Where("dept_name = ?", DeptName)
	}
	if DeptStatus != "" {
		curDb = Db.Where("dept_status = ?", DeptStatus)
	}
	curDb.Find(&sysDept)
	return sysDept
}

func GetSysDeptByName(deptName string) (sysDept entity.SysDept) {
	Db.Where("dept_name = ?", deptName).First(&sysDept)
	return sysDept
}

func CreateSysDept(sysDept entity.SysDept) bool {
	sysDeptByName := GetSysDeptByName(sysDept.DeptName)
	if sysDeptByName.ID > 0 {
		return false
	}
	// 1 = 公司
	if sysDept.DeptType == 1 {
		sysDept := entity.SysDept{
			DeptName:   sysDept.DeptName,
			DeptType:   sysDept.DeptType,
			DeptStatus: sysDept.DeptStatus,
			ParentId:   0,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Save(&sysDept)
		return true
	} else {
		sysDept := entity.SysDept{
			DeptName:   sysDept.DeptName,
			DeptType:   sysDept.DeptType,
			DeptStatus: sysDept.DeptStatus,
			ParentId:   sysDept.ParentId,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Save(&sysDept)
		return true
	}
}

func GetSysDepyById(id int) (sysDept entity.SysDept) {
	//Db.Where("id = ?", id).First(&sysDept)
	Db.First(&sysDept, id)
	return sysDept
}

func UpdateDysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	Db.First(&sysDept, dept.ID)
	sysDept.DeptType = dept.DeptType
	sysDept.DeptName = dept.DeptName
	sysDept.DeptStatus = dept.DeptStatus
	sysDept.ParentId = dept.ParentId
	Db.Save(&sysDept)
	return sysDept
}

func GetSysAdminDept(id int) (sysAdmin entity.SysAdmin) {
	Db.Where("id = ?", id).First(&sysAdmin)
	return sysAdmin
}

func DeleteSysDeptById(dto entity.SysDeptDto) bool {
	sysAdmin := GetSysAdminDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	Db.Delete("parent_id = ?", dto.Id).Delete(&entity.SysDept{})
	Db.Delete(&entity.SysDept{}, dto.Id)
	return true
}

func GetDeptList() (sysDeptVo []entity.SysDeptVo) {
	Db.Table("sys_dept").Select("id, parent_id, dept_name AS label").Scan(&sysDeptVo)
	return sysDeptVo
}

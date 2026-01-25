package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
	"go-admin/util"
	"time"
)

func GetSysMenuByName(menuName string) (menu entity.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&menu)
	return menu
}

func CreateSysMenu(menu entity.SysMenu) bool {
	sysMenuByName := GetSysMenuByName(menu.MenuName)
	if sysMenuByName.ID != 0 {
		return false
	}
	if menu.MenuType == 1 {
		sysMenu := entity.SysMenu{
			MenuName:   menu.MenuName,
			ParentId:   0,
			MenuType:   menu.MenuType,
			Icon:       menu.Icon,
			Sort:       menu.Sort,
			MenuStatus: menu.MenuStatus,
			Url:        menu.Url,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if menu.MenuType == 2 {
		sysMenu := entity.SysMenu{
			MenuName:   menu.MenuName,
			ParentId:   menu.ParentId,
			MenuType:   menu.MenuType,
			Icon:       menu.Icon,
			Sort:       menu.Sort,
			MenuStatus: menu.MenuStatus,
			Url:        menu.Url,
			CreateTime: util.HTime{Time: time.Now()},
			Value:      menu.Value,
		}
		Db.Create(&sysMenu)
		return true
	} else if menu.MenuType == 3 {
		sysMenu := entity.SysMenu{
			MenuName:   menu.MenuName,
			ParentId:   menu.ParentId,
			MenuType:   menu.MenuType,
			MenuStatus: menu.MenuStatus,
			Value:      menu.Value,
			Sort:       menu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

func GetSysMenuVoList() (sysMenu []entity.SysMenuVo) {
	Db.Table("sys_menu").Select("id, parent_id, menu_name AS label").Scan(&sysMenu)
	return sysMenu
}

func GetSysMenuById(menuId int) (menu entity.SysMenu) {
	//Db.Where("id = ?", menuId).First(&menu)
	Db.First(&menu, menuId)
	return menu
}

func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Sort = menu.Sort
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.ParentId = menu.ParentId
	sysMenu.Value = menu.Value
	sysMenu.MenuStatus = menu.MenuStatus
	Db.Save(&sysMenu)
	return sysMenu
}

func GetSysRoleMenu(id uint) (menu entity.SysRoleMenu) {
	Db.Where("menu_id = ?", id).First(&menu)
	return menu
}

func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	// 菜单已分配角色则不能删除
	sysRoleMenu := GetSysRoleMenu(dto.Id)
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	Db.Delete(entity.SysMenu{}, dto.Id)
	return true
}

func GetSysMenuList(menuName string, menuStatus string) (list []entity.SysMenu) {
	curDb := Db.Table("sys_menu").Order("sort")
	if menuName != "" {
		curDb = curDb.Where("menu_name = ?", menuName)
	}
	if menuStatus != "" {
		curDb = curDb.Where("menu_status = ?", menuStatus)
	}
	curDb.Find(&list)
	return list
}

func QueryMenuList(AdminId, MenuId uint) (menSvo []entity.MenuVo) {
	const status, menuStatus, menuType = 1, 2, 2
	Db.Table("sys_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").Scan(&menSvo)
	return menSvo
}

func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Order("sm.sort").Scan(&leftMenuVo)
	return leftMenuVo
}

// 登陆用户权限
func QueryUserPermissionList(Id uint) (valueVo []entity.ValueVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Not("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo)
	return valueVo
}

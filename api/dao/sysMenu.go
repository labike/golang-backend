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

package dao

// 具体的业务实现
import (
	"go-admin/api/entity"
	"go-admin/util"
	"time"
)
import . "go-admin/pkg/db"

// 岗位数据
func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

func CreateSysPost(sysPost entity.SysPost) bool {
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}
	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}
	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     sysPost.Remark,
	}
	tx := Db.Save(&addSysPost)
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

func GetSysPostList(PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, total int64) {
	curDb := Db.Table("sys_post")
	if PostName != "" {
		curDb = curDb.Where("post_name = ?", PostName)
	}
	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}
	if BeginTime != "" && BeginTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&total)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysPost)
	return sysPost, total
}

func GetPostById(id int) (sysPost entity.SysPost) {
	Db.Where("id = ?", id).First(&sysPost)
	return sysPost
}

func UpdateSysPost(post entity.SysPost) (sysPost entity.SysPost) {
	Db.First(&sysPost, post.ID)
	sysPost.PostName = post.PostName
	sysPost.PostStatus = post.PostStatus
	sysPost.PostCode = post.PostCode
	if post.Remark != "" {
		sysPost.Remark = post.Remark
	}
	Db.Save(&sysPost)
	return sysPost
}

func DeleteSysPostById(dto entity.SysPostIdDto) {
	Db.Delete(&entity.SysPost{}, dto.Id)
}

func BatchDeleteSysPost(dto entity.DelSysPostDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysPost{})
}

func UpdatePostStatus(dto entity.UpdateSysPostStatusDto) {
	var sysPost entity.SysPost
	Db.First(&sysPost, dto.Id)
	sysPost.PostStatus = dto.PostStatus
	Db.Save(&sysPost)
}

func SysPostVoList() (sysPostVo []entity.SysPostVo) {
	Db.Table("sys_post").Select("id, post_name").Scan(&sysPostVo)
	return sysPostVo
}

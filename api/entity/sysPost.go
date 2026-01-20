package entity

import "go-admin/util"

// 岗位模型(entity用于创建实体)
type SysPost struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	PostCode   string     `gorm:"column:post_code;varchar(64);comment:'岗位编码';NOT NULL" json:"postCode"`
	PostName   string     `gorm:"column:post_name;varchar(58);comment:'岗位名称';NOT NULL" json:"postName"`
	PostStatus int        `gorm:"column:post_status;default:1;comment:'状态(1->正常 2->停用)';NOT NULL" json:"postStatus"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
	Remark     string     `gorm:"column:remark;varchar(500);comment:'备注';NOT NULL" json:"remark"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

type SysPostIdDto struct {
	Id uint `json:"id"`
}

type DelSysPostDto struct {
	Ids []uint
}

type UpdateSysPostStatusDto struct {
	Id         uint
	PostStatus int
}

// 岗位下拉列表
type SysPostVo struct {
	Id       uint   `json:"id"`
	PostName string `json:"postName"`
}

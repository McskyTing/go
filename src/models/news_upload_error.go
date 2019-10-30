package models

import (
	"time"
)

type NewsUploadError struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Uid        string    `xorm:"not null default '' comment('新闻全局uid') index VARCHAR(32)"`
	Status     int       `xorm:"not null default 0 comment('上传状态0上传失败1上传成功') TINYINT(4)"`
	CreateTime time.Time `xorm:"not null comment('记录创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null comment('记录更新时间') DATETIME"`
}

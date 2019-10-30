package models

import (
	"time"
)

type InfoDetail404 struct {
	Id         int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	Uid        string    `xorm:"not null default '' comment('新闻全局唯一ID') index VARCHAR(32)"`
	InfoId     string    `xorm:"not null default '' comment('资讯ID') index VARCHAR(50)"`
	Sites      int       `xorm:"not null default 0 comment('网站来源，1-东方资讯；2-百度 ...') TINYINT(4)"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') index DATETIME"`
}

package models

import (
	"time"
)

type News404 struct {
	Id         int       `xorm:"not null pk autoincr comment('ID') INT(10)"`
	NewsId     string    `xorm:"not null comment('新闻全局唯一ID') unique VARCHAR(42)"`
	Status     int       `xorm:"not null default 1 comment('post状态，1-待发送，0-已发送') TINYINT(1)"`
	CreateTime time.Time `xorm:"not null comment('添加时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null comment('修改时间') DATETIME"`
}

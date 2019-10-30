package models

import (
	"time"
)

type NewsErrorUrl struct {
	Id         int       `xorm:"not null pk autoincr comment('自增id') INT(10)"`
	Url        string    `xorm:"not null default '' comment('抓取失败地址') VARCHAR(255)"`
	ErrorType  int       `xorm:"not null default 0 comment('错误类型1 数据库错误  2 抓取错误') TINYINT(2)"`
	CreateTime time.Time `xorm:"not null comment('添加时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME"`
	Source     int       `xorm:"not null default 0 comment('抓取所在表') TINYINT(4)"`
	Status     int       `xorm:"not null default 0 comment('再次抓取是否成功') TINYINT(4)"`
}

package models

import (
	"time"
)

type SnHostConfig struct {
	Id          int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	ProductName string    `xorm:"not null default '' comment('产品名称') VARCHAR(100)"`
	ProductId   int       `xorm:"not null default 100000 comment('产品id') INT(11)"`
	HostName    string    `xorm:"not null default '' comment('宿主名称') VARCHAR(100)"`
	HostId      int       `xorm:"not null comment('宿主id = 产品id*1000+maxhost_id+1') INT(11)"`
	IsDel       int       `xorm:"not null default 0 comment('是否被删除0未删除1删除') TINYINT(1)"`
	IsPublish   int       `xorm:"not null default 0 comment('是否上架0未上架1上架') TINYINT(1)"`
	CreateTime  time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null comment('修改时间') DATETIME"`
}

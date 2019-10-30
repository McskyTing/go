package models

import (
	"time"
)

type CrawlPartnerConfig struct {
	Id          int       `xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	IsDel       int       `xorm:"not null default 0 comment('是否被删除0未删除1删除') TINYINT(1)"`
	CreateTime  time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null comment('修改时间') DATETIME"`
	PartnerName string    `xorm:"not null default '' comment('合作方中文名称') VARCHAR(64)"`
	Partner     string    `xorm:"not null default '' comment('合作方英文名称(字母数字下划线)') VARCHAR(64)"`
	JoinTime    time.Time `xorm:"not null comment('接入时间') DATETIME"`
	Num         int       `xorm:"not null default 0 comment('新闻量级(条/天)') INT(11)"`
	Remark      string    `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	Appid       string    `xorm:"not null default '' comment('appid uniqueid生成') VARCHAR(20)"`
	Appsrcret   string    `xorm:"not null default '' comment('appsrcret') VARCHAR(64)"`
	Status      int       `xorm:"not null default 0 comment('状态0未启用1启用') TINYINT(4)"`
}

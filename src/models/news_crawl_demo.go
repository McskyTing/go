package models

import (
	"time"
)

type NewsCrawlDemo struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ItemId       string    `xorm:"not null default '' comment('资讯id') unique VARCHAR(50)"`
	Title        string    `xorm:"not null default '' comment('标题') VARCHAR(200)"`
	Source       string    `xorm:"not null default '' comment('来源') VARCHAR(50)"`
	HotTime      time.Time `xorm:"not null comment('成为热门时间') DATETIME"`
	DeleteStatus int       `xorm:"not null default 0 comment('是否删除0-不删除1-删除') TINYINT(1)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

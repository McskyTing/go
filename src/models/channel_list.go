package models

import (
	"time"
)

type ChannelList struct {
	Id                      int       `xorm:"not null pk autoincr comment('自增长id') INT(11)"`
	StandardChannel         string    `xorm:"not null default '' comment('标准channel') VARCHAR(50)"`
	StandardChannelAlphabet string    `xorm:"not null comment('标准channel对应的拼音') VARCHAR(50)"`
	ExternalChannel         string    `xorm:"not null default '' comment('外部channel') unique(uix_et) VARCHAR(50)"`
	ThirdSource             string    `xorm:"not null default '' comment('内容来源') unique(uix_et) VARCHAR(20)"`
	CreateTime              time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('插入时间') DATETIME"`
	UpdateTime              time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

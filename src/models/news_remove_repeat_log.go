package models

import (
	"time"
)

type NewsRemoveRepeatLog struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	NewsDate  time.Time `xorm:"not null comment('去重新闻日期') DATE"`
	StartHour int       `xorm:"not null comment('查询起始小时') TINYINT(4)"`
	EndHour   int       `xorm:"not null comment('查询结束小时') TINYINT(4)"`
	StartId   int64     `xorm:"not null comment('查询最小id') BIGINT(20)"`
	Status    int       `xorm:"not null default 0 comment('去重后数据是否入正式表') TINYINT(1)"`
}

package models

import (
	"time"
)

type Hotwords struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Keywords   string    `xorm:"not null default '' comment('抓回来的关键词') unique VARCHAR(100)"`
	Searches   int       `xorm:"not null default 0 comment('接口返回数据,搜索次数') index INT(10)"`
	ChangeRate int       `xorm:"comment('对应接口返回字段') INT(10)"`
	IsNew      int       `xorm:"comment('对应接口返回字段') TINYINT(3)"`
	Trend      string    `xorm:"comment('对应接口返回字段') VARCHAR(20)"`
	Gentime    int       `xorm:"not null default 0 comment('抓取接口返回字段,用于排序') index INT(10)"`
	DataFrom   string    `xorm:"not null default '' comment('热词从哪里抓取过来的') VARCHAR(20)"`
	CreatedAt  time.Time `xorm:"not null comment('插入数据库时间') DATETIME"`
	UpdatedAt  time.Time `xorm:"comment('热词更新时间, 默认与创建时间相同') DATETIME"`
}

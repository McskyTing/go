package models

import (
	"time"
)

type NewsRepeat struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	BaseNewsId    string    `xorm:"not null default '' comment('基准id') VARCHAR(38)"`
	SimilarNewsId string    `xorm:"not null default '' comment('相似id') index VARCHAR(38)"`
	Score         int       `xorm:"not null comment('相似度') TINYINT(4)"`
	DeleteStatus  int       `xorm:"not null default 0 comment('是否删除0-不删除1-删除') TINYINT(1)"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') index DATETIME"`
	UpdateTime    time.Time `xorm:"not null comment('更新时间') DATETIME"`
}

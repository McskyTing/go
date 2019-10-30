package models

import (
	"time"
)

type DelYidianNews struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	InfoId     string    `xorm:"not null default '' VARCHAR(20)"`
	NewsId     string    `xorm:"not null default '' comment('唯一id') unique VARCHAR(100)"`
	DeleteTime time.Time `xorm:"not null default '0000-00-00 00:00:00' index DATETIME"`
	IsUpdated  int       `xorm:"not null default 0 comment('是否已经同步删除 1是') index TINYINT(4)"`
}

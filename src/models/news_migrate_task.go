package models

import (
	"time"
)

type NewsMigrateTask struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	RequestEndDate string    `xorm:"not null comment('去重请求最大时间格式YYMMDDHH') unique VARCHAR(10)"`
	RequestStartId int64     `xorm:"not null default 0 comment('去重请求起始id即上一次请求最大id') BIGINT(20)"`
	RepeatMaxId    int64     `xorm:"not null default 0 comment('本次去重结果最大id') BIGINT(20)"`
	RepeatCount    int       `xorm:"not null comment('去重结果数量') INT(11)"`
	MigrateStartId int64     `xorm:"not null default 0 comment('本次迁移记录起始id') BIGINT(20)"`
	MigrateCount   int       `xorm:"not null default 0 comment('待迁移数据总量') INT(11)"`
	TaskStatus     int       `xorm:"not null default 0 comment('任务状态0去重接口错误1接口成功迁移未完成2迁移已完成') TINYINT(4)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('任务创建时间') DATETIME"`
	UpdateTime     time.Time `xorm:"not null comment('任务更新时间') DATETIME"`
}

package models

import (
	"time"
)

type SensitiveWords struct {
	Id             int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	SensitiveWords string    `xorm:"not null comment('敏感词,每个敏感词用中文分号、') TEXT"`
	Type           string    `xorm:"not null comment('风险类别') VARCHAR(50)"`
	AuditLevel     int       `xorm:"not null default 1 comment('审核级别  1=初审  2=复审  3=终审') TINYINT(1)"`
	TotalWords     int       `xorm:"not null default 0 comment('总敏感词数') INT(10)"`
	HostId         string    `xorm:"not null comment('宿主ID 多个逗号分隔') VARCHAR(500)"`
	Programme      int       `xorm:"not null default 1 comment('命中后处理方案  1=直接屏蔽下发  2=加入新闻库待审核') TINYINT(1)"`
	Operator       string    `xorm:"not null comment('操作人') VARCHAR(50)"`
	CreateTime     time.Time `xorm:"not null comment('创建时间') TIMESTAMP"`
	UpdateTime     time.Time `xorm:"not null comment('更新时间') TIMESTAMP"`
}

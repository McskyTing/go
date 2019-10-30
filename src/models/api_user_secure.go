package models

import (
	"time"
)

type ApiUserSecure struct {
	Id              int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	UserName        string    `xorm:"not null comment('合作方名称') VARCHAR(50)"`
	SourceName      string    `xorm:"not null default '' comment('第三方源信息，third_source') unique VARCHAR(50)"`
	AppId           string    `xorm:"not null default '' comment('用户ID，唯一，随机生成20位字符串') unique VARCHAR(20)"`
	AppSecret       string    `xorm:"not null default '' comment('用户密钥，唯一，随机生成50位字符串') VARCHAR(50)"`
	NewsLevel       int       `xorm:"not null default 0 comment('新闻量级') INT(11)"`
	Remarks         string    `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	CooperateStatus int       `xorm:"not null default 0 comment('合作方状态，0是未开启，1是开启') TINYINT(1)"`
	CooperateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('合作时间') DATETIME"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

package models

import (
	"time"
)

type SensitiveLog struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	Type        int       `xorm:"not null default 0 comment('错误类型 0代表请求日志，1代表敏感词日志') TINYINT(3)"`
	Word        string    `xorm:"not null default '' comment('敏感词') VARCHAR(250)"`
	Title       string    `xorm:"not null default '' comment('标题') VARCHAR(500)"`
	ThirdSource string    `xorm:"not null default '' comment('内容来源:df_server-东方头条,baidu-百度feeds,wuli-唔哩头条,yidian_news => 一点资讯新闻,yidian_video-一点资讯视频,rabbit-好兔短视频,weishi-微视小视频') VARCHAR(20)"`
	HostId      int       `xorm:"not null default 0 comment('宿主id') INT(11)"`
	Url         string    `xorm:"not null default '' comment('请求的地址') VARCHAR(2000)"`
	Num         int       `xorm:"not null comment('新闻下发条数，type为0值才有意义') TINYINT(3)"`
	CreateAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
}

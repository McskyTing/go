package models

import (
	"time"
)

type NewsCheckLog201907 struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	InfoId      string    `xorm:"not null comment('资讯ID，来自于第三方新闻源') VARCHAR(255)"`
	ThirdSource string    `xorm:"not null comment('内容来源:df_server-东方头条,baidu-百度feeds,wuli-唔哩头条,yidian_news => 一点资讯新闻,yidian_video-一点资讯视频,rabbit-好兔短视频,weishi-微视小视频') VARCHAR(255)"`
	UrlSource   int       `xorm:"not null comment('链接来源,0:客户端，1:资讯列表,2:全表') TINYINT(1)"`
	Url         string    `xorm:"not null comment('资讯链接') VARCHAR(2000)"`
	OriginUrl   string    `xorm:"not null comment('客户端上报的原始url') VARCHAR(2000)"`
	Uid         int       `xorm:"not null comment('passid') INT(11)"`
	HttpCode    int       `xorm:"not null comment('http code') SMALLINT(4)"`
	ErrorCode   int       `xorm:"not null comment('真实错误码') SMALLINT(255)"`
	Imei        string    `xorm:"not null comment('imei') VARCHAR(20)"`
	UserInfo    string    `xorm:"not null comment('用户信息') VARCHAR(2000)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

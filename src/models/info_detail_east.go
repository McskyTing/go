package models

import (
	"time"
)

type InfoDetailEast struct {
	Id           int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	Uid          string    `xorm:"not null comment('新闻全局唯一ID') unique VARCHAR(32)"`
	Title        string    `xorm:"not null default '' comment('资讯标题') index VARCHAR(255)"`
	InfoId       string    `xorm:"not null default '' comment('资讯ID') unique VARCHAR(50)"`
	InfoUrl      string    `xorm:"not null default '' comment('资讯链接') index VARCHAR(512)"`
	HttpsUrl     string    `xorm:"not null default '' comment('资讯https链接') VARCHAR(5000)"`
	LargePic     string    `xorm:"not null default '' comment('大图链接列表') VARCHAR(2048)"`
	MiniPic      string    `xorm:"not null default '' comment('小图链接列表') VARCHAR(2048)"`
	MiniImgSize  int       `xorm:"not null default 0 comment('小图数量') INT(11)"`
	CommentCount int       `xorm:"not null default 0 comment('评论数') INT(11)"`
	Source       string    `xorm:"not null default '' comment('资讯来源') index VARCHAR(50)"`
	ClickRate    string    `xorm:"not null default 0.0000 comment('点击率') DECIMAL(10,4)"`
	HotNews      int       `xorm:"not null default 0 comment('是否是热点（0：非热点; 1：热点）') TINYINT(4)"`
	IsRecommend  int       `xorm:"not null default 0 comment('是否是推荐（0：非推荐; 1：推荐）') TINYINT(4)"`
	IsVideo      int       `xorm:"not null default 0 comment('是否是视频新闻（0：不是视频新闻;1:是视频新闻）') TINYINT(4)"`
	VideoTime    int       `xorm:"not null default 0 comment('视频时间(单位：ms)') INT(11)"`
	Type         string    `xorm:"not null default '' comment('资讯类别，如：toutiao（头条）、shehui（社会）、guonei（国内）...') index VARCHAR(50)"`
	RowKey       string    `xorm:"not null default '' comment('rowkey') VARCHAR(50)"`
	BigPic       int       `xorm:"not null default 0 comment('bigpic') TINYINT(4)"`
	Content      string    `xorm:"not null comment('资讯内容') TEXT"`
	IsApi        int       `xorm:"not null default 0 comment('来源api接口还是爬虫，1-api；2-爬虫') TINYINT(4)"`
	Sites        int       `xorm:"not null default 0 comment('网站来源，1-东方资讯；2-百度 ...') TINYINT(4)"`
	PublishDate  time.Time `xorm:"not null comment('资讯发布时间') index DATETIME"`
	CreateTime   time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null comment('更新时间') index DATETIME"`
}

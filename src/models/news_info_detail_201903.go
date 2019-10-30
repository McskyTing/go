package models

import (
	"time"
)

type NewsInfoDetail201903 struct {
	Id           int64     `xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	NewsId       string    `xorm:"not null comment('新闻全局唯一ID') unique VARCHAR(42)"`
	InfoId       string    `xorm:"not null default '' comment('资讯ID') unique(uix_info_id) VARCHAR(50)"`
	ThirdSource  string    `xorm:"not null default '' comment('内容来源:df_server-东方头条,baidu-百度feeds,wuli-唔哩头条,yidian_news => 一点资讯新闻,yidian_video-一点资讯视频,rabbit-好兔短视频,weishi-微视小视频') unique(uix_info_id) VARCHAR(20)"`
	ContentType  string    `xorm:"not null default '' comment('内容类型：news-资讯,short_video-短视频,mini_video-小视频') VARCHAR(20)"`
	PublishDate  time.Time `xorm:"not null comment('资讯发布时间') index DATETIME"`
	UpdateAt     time.Time `xorm:"not null comment('资讯更新时间') DATETIME"`
	Title        string    `xorm:"not null default '' comment('资讯标题') index VARCHAR(255)"`
	Source       string    `xorm:"not null default '' comment('资讯来源') VARCHAR(50)"`
	Nickname     string    `xorm:"not null default '' comment('作者名称') VARCHAR(50)"`
	InfoUrl      string    `xorm:"not null default '' comment('资讯链接') index VARCHAR(2000)"`
	HttpsUrl     string    `xorm:"not null default '' comment('资讯https链接') VARCHAR(2000)"`
	PureUrl      string    `xorm:"not null default '' comment('纯净版链接') VARCHAR(1000)"`
	NewsTag      string    `xorm:"not null default '' comment('新闻标签') VARCHAR(255)"`
	Category     string    `xorm:"not null default '' comment('资讯分类') VARCHAR(255)"`
	Type         string    `xorm:"not null default '' comment('所属频道，如：toutiao（头条）、shehui（社会）、guonei（国内）...') VARCHAR(50)"`
	Channel      string    `xorm:"not null default '' comment('标准频道，自定义频道') VARCHAR(50)"`
	Summary      string    `xorm:"not null default '' comment('资讯摘要') VARCHAR(255)"`
	Content      string    `xorm:"not null comment('资讯内容') TEXT"`
	BigPic       int       `xorm:"not null default 0 comment('是否是大图1大图0不是大图') TINYINT(1)"`
	LargePic     string    `xorm:"not null default '' comment('大图链接列表') VARCHAR(2048)"`
	MiniPic      string    `xorm:"not null default '' comment('小图链接列表') VARCHAR(2048)"`
	IsPure       int       `xorm:"not null comment('是否纯净链接  1-是  0-不是') TINYINT(1)"`
	MiniImgSize  int       `xorm:"not null default 0 comment('小图数量') INT(11)"`
	CommentCount int       `xorm:"not null default 0 comment('评论数') INT(11)"`
	LikeCount    int       `xorm:"not null default 0 comment('点赞数量') INT(10)"`
	ShareCount   int       `xorm:"not null default 0 comment('分享数量') INT(10)"`
	CollectCount int       `xorm:"not null default 0 comment('收藏数量') INT(10)"`
	DislikeCount int       `xorm:"not null default 0 comment('不喜欢/踩数量') INT(10)"`
	HotNews      int       `xorm:"not null default 0 comment('是否是热点（0：非热点; 1：热点）') TINYINT(4)"`
	IsRecommend  int       `xorm:"not null default 0 comment('是否是推荐（0：非推荐; 1：推荐）') TINYINT(4)"`
	IsVideo      int       `xorm:"not null default 0 comment('是否是视频新闻（0：不是视频新闻;1:是视频新闻）') TINYINT(4)"`
	IsChoice     int       `xorm:"not null default 0 comment('是否精选0-不是，1-是') TINYINT(4)"`
	VideoUrl     string    `xorm:"not null default '' comment('视频源地址') VARCHAR(2000)"`
	VideoTime    int       `xorm:"not null default 0 comment('视频时间(单位：ms)') INT(11)"`
	PlayCount    int       `xorm:"not null default 0 comment('播放数量') INT(11)"`
	IsRepeat     int       `xorm:"not null default 0 comment('是否重复0-不重复，1-重复') TINYINT(1)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') index DATETIME"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') index DATETIME"`
}

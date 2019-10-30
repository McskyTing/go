package models

import (
	"time"
)

type SnLogData struct {
	Id                 int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	ProductId          int       `xorm:"not null default 0 comment('产品id') INT(11)"`
	HostId             int       `xorm:"not null default 0 comment('宿主id ') INT(11)"`
	MediaId            int       `xorm:"not null default 0 comment('媒体位id') INT(11)"`
	AppVersion         string    `xorm:"not null default '' comment('各宿主对应的版本号') VARCHAR(50)"`
	ChannelHeader      string    `xorm:"not null default '' comment('渠道号') VARCHAR(200)"`
	ContentType        string    `xorm:"not null default '' comment('内容类型') VARCHAR(30)"`
	ChannelType        string    `xorm:"not null default '' comment('频道') VARCHAR(30)"`
	ChargeId           string    `xorm:"not null default '' comment('第三方计费id') VARCHAR(50)"`
	EditType           string    `xorm:"not null default '' comment('运营新闻类别') VARCHAR(30)"`
	SdkVersion         string    `xorm:"not null default '' comment('SDK对应的版本号') VARCHAR(30)"`
	RecommendType      string    `xorm:"not null default '' comment('算法') VARCHAR(30)"`
	SdkLiveNum         int       `xorm:"not null default 0 comment('SDK活跃用户数') INT(11)"`
	AvgClickAddRate    string    `xorm:"not null default 0.00 comment('人均点击数周增幅') DECIMAL(10,2)"`
	AvgClickNum        string    `xorm:"not null default 0.00 comment('人均点击数') DECIMAL(10,2)"`
	AvgReadAddRate     string    `xorm:"not null default 0.00 comment('人均阅读时长周增幅') DECIMAL(10,2)"`
	AvgReadTime        string    `xorm:"not null default 0.00 comment('人均阅读时长/毫秒') DECIMAL(10,2)"`
	CtrAddRate         string    `xorm:"not null default 0.00 comment('CTR周增幅') DECIMAL(10,2)"`
	CtrClickRate       string    `xorm:"not null default 0.00 comment('CTR点击率%') DECIMAL(10,2)"`
	ReadAddRate        string    `xorm:"not null default 0.00 comment('深度周增幅') DECIMAL(10,2)"`
	ReadDepth          string    `xorm:"not null default 0.00 comment('阅读深度') DECIMAL(10,2)"`
	ExposureRate       string    `xorm:"not null default 0.00 comment('曝光用户占比%') DECIMAL(10,2)"`
	ClickRate          string    `xorm:"not null default 0.00 comment('点击用户占比%') DECIMAL(10,2)"`
	ClickAvgReadTime   string    `xorm:"not null default 0.00 comment('点击用户人均阅读时长/毫秒') DECIMAL(10,2)"`
	ExposurePv         int       `xorm:"not null default 0 comment('曝光PV') INT(11)"`
	ExposureUv         int       `xorm:"not null default 0 comment('曝光uv') INT(11)"`
	ClickPv            int       `xorm:"not null default 0 comment('点击PV') INT(11)"`
	ClickUv            int       `xorm:"not null default 0 comment('点击uv') INT(11)"`
	Day                time.Time `xorm:"comment('日志日期') index DATE"`
	EveryClickReadTime string    `xorm:"not null default 0.00 comment('单次点击阅读时长') DECIMAL(10,2)"`
	StayTime           int       `xorm:"not null default 0 comment('停留时长 /毫秒') INT(11)"`
}

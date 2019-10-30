package models

import (
	"time"
)

type SnMbNewsLog20190916 struct {
	Id             int       `xorm:"not null pk autoincr comment('ID') INT(10)"`
	HostId         int       `xorm:"not null default 0 comment('宿主id') INT(10)"`
	AppVersion     string    `xorm:"not null default '' comment('app版本号') VARCHAR(20)"`
	Imei           string    `xorm:"not null default '' comment('国际移动设备身份码') VARCHAR(64)"`
	Mac            string    `xorm:"not null default '' comment('设备mac地址') VARCHAR(64)"`
	SdkChannel     string    `xorm:"not null default '' comment('SDK渠道号') VARCHAR(64)"`
	Package        string    `xorm:"not null default '' comment('宿主包名') VARCHAR(64)"`
	Passid         string    `xorm:"not null default '' comment('用户中心账号') VARCHAR(32)"`
	Uid            string    `xorm:"not null default '' comment('用户id') VARCHAR(64)"`
	SdkVersion     string    `xorm:"not null default '' comment('sdk版本号') VARCHAR(32)"`
	SdkVersionCode int       `xorm:"not null default 0 comment('SDK内部版本号') INT(10)"`
	ChannelType    string    `xorm:"not null default '' comment('频道类型') index VARCHAR(32)"`
	EventType      int       `xorm:"not null default 0 comment('事件类型') TINYINT(1)"`
	BusinessId     int       `xorm:"not null default 0 comment('业务类型') INT(11)"`
	SceneId        int       `xorm:"not null default 0 comment('业务场景,针对研究院服务的不同项目') INT(10)"`
	ItemType       string    `xorm:"not null default '' comment('业务类型') VARCHAR(64)"`
	Url            string    `xorm:"not null default '' comment('资讯,视频地址') VARCHAR(1024)"`
	StayTime       int       `xorm:"not null default 0 comment('停留时长毫秒') INT(10)"`
	InfoId         string    `xorm:"not null default '' comment('匹配url后的id') VARCHAR(64)"`
	NewsId         string    `xorm:"not null default '' comment('新闻id') VARCHAR(64)"`
	ContentType    string    `xorm:"not null default '' comment('新闻内容类型') index VARCHAR(64)"`
	RecommendType  string    `xorm:"not null default '' comment('算法类型') index VARCHAR(64)"`
	EditType       string    `xorm:"not null default '' comment('运营新闻类别') index VARCHAR(64)"`
	Ip             string    `xorm:"not null default '' comment('用户ip') VARCHAR(20)"`
	Time           int       `xorm:"not null default 0 comment('上报时间戳') INT(10)"`
	Day            time.Time `xorm:"comment('日期') index DATE"`
	Title          string    `xorm:"not null default '' comment('新闻title') VARCHAR(1024)"`
	ChannelHeader  string    `xorm:"not null default '' comment('渠道号') VARCHAR(50)"`
	MediaId        int       `xorm:"not null default 0 comment('媒体位id') INT(11)"`
	ChargeId       string    `xorm:"not null default '' comment('第三方计费id') VARCHAR(50)"`
	PlayTime       int       `xorm:"not null default 0 comment('小视频播放时长') INT(10)"`
	Request        string    `xorm:"not null default '' comment('容错逻辑') VARCHAR(10)"`
	Direction      string    `xorm:"not null default '' comment('刷新方向') VARCHAR(10)"`
	RefreshType    string    `xorm:"not null default '' comment('刷新类型') VARCHAR(10)"`
	Redpacket      int       `xorm:"not null default 0 comment('是否红包,0否1是') TINYINT(2)"`
	NewsIdFrom     string    `xorm:"not null default '' comment('相关搜索词曝光时，当前新闻id') VARCHAR(64)"`
	BigDataSource  string    `xorm:"not null default '' comment('大数据推荐源') VARCHAR(64)"`
	LocationType   string    `xorm:"not null default '' comment('下发新闻类型') VARCHAR(64)"`
	IsPure         int       `xorm:"not null default 0 comment('下发新闻类型') TINYINT(2)"`
	ThirdQid       string    `xorm:"not null default '' comment('第三方计费链') VARCHAR(64)"`
}

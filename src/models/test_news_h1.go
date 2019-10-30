package models

type TestNewsH1 struct {
	Id     int64  `xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	NewsId string `xorm:"not null comment('新闻全局唯一ID') VARCHAR(42)"`
}

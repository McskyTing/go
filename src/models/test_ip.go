package models

type TestIp struct {
	Ip string `xorm:"default '' VARCHAR(20)"`
}

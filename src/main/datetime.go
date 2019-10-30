package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	//时间转化为string
	//在go语言里将日期格式化并不是yyyy-MM-dd HH:ss:mm 而是用"2006-01-02 15:04:05具体含义如下
	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Println("日期类型当前时间: ", now)
	fmt.Println("字符串类型当前时间: ", timeStr)
	//string转化为时间
	date, _ := time.Parse("2006-01-02 15:04:05", "2017-08-29 08:37:18")
	fmt.Println("string转日期:", date)
	//判断两个时间先后
	trueOrFalse := date.After(now)
	if trueOrFalse == true {
		fmt.Println("2017-08-29 08:37:18在", timeStr, "之后")
	} else {
		fmt.Println("2017-08-29 08:37:18在", timeStr, "之前")
	}
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// -代表时间之前,不加代表时间之后 m表示分钟,如20分钟前
	m, _ := time.ParseDuration("-20m")
	m1 := now.Add(m)
	fmt.Println("10分钟之前：", m1)
	// h代表小时 一天之前为-24h
	h, _ := time.ParseDuration("-8h")
	h1 := now.Add(h)
	fmt.Println("8小时之前", h1)
	//	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)
	//计算两个时间差几秒
	sec := now.Sub(m1)
	fmt.Println(sec.Seconds(), "秒")
	//计算两个时间差几分钟
	minutes := now.Sub(m1)
	fmt.Println(minutes.Minutes(), "分钟")
	//计算两个时间差几小时
	hours := now.Sub(h1)
	fmt.Println(hours.Hours(), "小时")
	//计算两个时间差几天
	day := now.Sub(d1)
	fmt.Println(day.Hours()/24, "天")
	//注意:splite3数据库中字段如果是datetime类型获取数据时格式转换会有问题
	//如2017-08-29 08:37:18这样的时间从数据库获取后会变成2017-08-29T08:37:18Z
	//进行格式转化之后不能比较，所以需要将T和Z替换为" "
	//不知道其他数据库有没有这样的问题
}

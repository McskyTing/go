package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int到string
	str := strconv.Itoa(int(1))
	fmt.Println("int转string", str)
	//int64到string
	str2 := strconv.FormatInt(int64(1), 10)
	fmt.Println("int64转string", str2)
	//float64转string formatFloat只能接收float64如果想用float32需要强转float64(float32(0.8))
	//下面是参数说明
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	str3 := strconv.FormatFloat(float64(0.8), 'f', -1, 32)
	fmt.Println("float32转string", str3)
	//string到int  有异常的都不进行处理这个后面说
	i, _ := strconv.Atoi("10")
	fmt.Println("strin转int", i)
	//string 到int64
	i64, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println("strin转int64", i64)
	//string转float64 如果想转float32 用float32(fl32)强转一下就可以
	fl32, _ := strconv.ParseFloat("3.1415926535", 32/64)
	fmt.Println("strin转float64", fl32)
}

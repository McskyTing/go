package main

import (
	"log"
	"strconv"
)

//捕获因未知输入导致的程序异常
func catch(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	return nums[1] * nums[2] * nums[3] //index out of range
}

//主动抛出 panic，不推荐使用，可能会导致性能问题
func toFloat64(num string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[W]", r)
		}
	}()
	if num == "" {
		panic("param is null") //主动抛出 panic
	}
	return strconv.ParseFloat(num, 10)
}
func main() {
	catch(2, 8)
	toFloat64("")
}

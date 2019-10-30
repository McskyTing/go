package main

import (
	"fmt"
)

func main() {
	fmt.Println("start================")
	fmt.Println("t5返回值", t5())
	fmt.Println("end================")
}
func t1() {
	fmt.Println("t1")
}
func t2() {
	fmt.Println("t2")
}
func t3() {
	fmt.Println("t3")
}
func t4() {
	fmt.Println("t4")
}
func t5() int {
	defer t1()
	defer t2()
	defer t3()
	defer t4()
	return 1
}

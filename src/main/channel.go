package main

import "fmt"

func fibonacci1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Println(cap(c))
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)
}

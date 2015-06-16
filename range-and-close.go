package main

import (
	"fmt"
)

func fibonacci(c chan int) {
	x, y := 0, 1
	for i := 0; i < cap(c); i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(c)
	for i := range c {
		fmt.Println(i)
	}
}

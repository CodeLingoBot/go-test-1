package main

import (
	"fmt"
	"sync"
)

func main() {

	var x, y int
	var loadSync sync.Once
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		x = 1
		fmt.Println("y:", y)
		ch1 <- struct{}{}
	}()
	go func() {
		y = 1
		fmt.Println("x:", x)
		ch2 <- struct{}{}
	}()
	<-ch1
	<-ch2
}

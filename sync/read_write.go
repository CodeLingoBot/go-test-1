package main

import (
	"fmt"
)

func read(data <-chan int) {
	for d := range data {
		fmt.Printf("%d ", d)
	}
}
func write(d int, data chan<- int) bool {
	select {
	case data <- d:
		return true
	default:
		return false
	}
}

func main() {
	const size int = 50
	data := make(chan int, 5)
	var x [size]int
	for i := 0; i < size; i++ {
		x[i] = i
	}
	go read(data)
	for i := 0; i < size; i++ {
		//fmt.Printf("%t ", write(i, data))
		write(i, data)
		//data <- i
	}
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

var pc [256]byte

func initPc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func print() {
	for _, v := range pc {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	initPc()
	print()
	fmt.Println(GetPrice(1.11))
}

package main

import "fmt"

type Flags uint

const (
	FlagUp Flags=1<<iota
	FlagBoradcast
	FlagLoopback
)

const (
	_=1<<(10*iota)
	KiB
	MiB
	GiB
	TiB
)

func output(){
	fmt.Println(FlagUp)
	fmt.Println(FlagBoradcast)
	fmt.Println(FlagLoopback)
}

func outputUnit(){
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
}

func main(){
	output()
	outputUnit()
}
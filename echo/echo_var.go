package main

var Label string

type Price float64

func GetPrice(p float64) Price {
	return Price(p)
}

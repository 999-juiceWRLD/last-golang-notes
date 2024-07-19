package main

import (
	"fmt"
)

const (
	ETH = iota
	WAN
)

func main() {
	rates := [3]float64{
		0: 0.5,
		1: 2.5,
		2: 1.5,
	}

	fmt.Println(rates)
	fmt.Println(rates[0])
	fmt.Println(rates[1])
	fmt.Println(rates[2])

	cryptoRates := [3]float64{
		ETH: 25.5,
		WAN: 120.5,
	}

	fmt.Printf("1 BTC is %g ETH\n", cryptoRates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", cryptoRates[WAN])
}
package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active ?:", s.active)
}

func defaultSubsciber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subsciber1 := defaultSubsciber("Amain singh")
	subsciber1.rate = 4.99
	printInfo(subsciber1)

	subsciber2 := defaultSubsciber("Beth Rayn")
	printInfo(subsciber2)
}

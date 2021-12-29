package main

import "fmt"

type subscriber2 struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber2) {
	s.rate = 4.99
}

func main() {
	var s subscriber2
	applyDiscount(&s)
	fmt.Println(s.rate)
}

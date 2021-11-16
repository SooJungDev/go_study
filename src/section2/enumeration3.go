package main

import "fmt"

func main() {

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("D:", DEFAULT)
	fmt.Println("S:", SILVER)
	fmt.Println("G:", GOLD)
	fmt.Println("P:", PLATINUM)
}

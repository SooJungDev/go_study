package main

import (
	"fmt"
	"section4/lib"
)

func init() {
	fmt.Println("Init1 Method start")
}

func init() {
	fmt.Println("Init2 Method start")
}

func init() {
	fmt.Println("Init3 Method start")
}

func init() {
	fmt.Println("Init4 Method start")
}

func main() {
	lib.CheckNum(10)
	fmt.Println("Main Method start")
}

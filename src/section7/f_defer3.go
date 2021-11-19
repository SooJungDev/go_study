package main

import "fmt"

func main() {
	//예제1
	stack()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	}
}

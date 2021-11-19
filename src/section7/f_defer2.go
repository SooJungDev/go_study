package main

import (
	"fmt"
)

func main() {
	//예제1
	sayHello("Golang!")
}

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi")
	}()
}

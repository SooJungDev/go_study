package main

import "fmt"

/**
  구조체체*/

func main() {

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

}

package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활용
	//예제1 type 구조체명 타입
	car1 := struct {
		name, color string
	}{"520d", "red"}

	fmt.Println("ex1 : ", car1)
	fmt.Println("ex1 : %#v1\n", car1)

	//예제2
	cars := []struct{ name, color string }{{"520d", "red"}, {"520i", "white"}, {"528i", "blue"}}

	for _, c := range cars {
		fmt.Printf("%s, %s ------(%#v)\n", c.name, c.color, c)
	}
}

package main

import "fmt"

func main() {
	//예제 1
	i := 7
	p := &i

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p++

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p = 7777 // 포인터 변수 역 참조 값 변경
	fmt.Println("ex1 : ", i, *p, &i, p)

	i = 77
	fmt.Println("ex1 : ", i, *p, &i, p)
}

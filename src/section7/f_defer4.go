package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func a() {
	defer end(start("b")) // 중첩 함수 주의!!! end 함수만 마지막에 실행됨!!
	fmt.Println("in a")
}

func end(t string) {
	fmt.Println("end :", t)
}

func main() {
	// 예제1
	a()
}

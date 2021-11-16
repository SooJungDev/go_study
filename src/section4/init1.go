package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init Method start")
}

func main() {
	// init: 패키지 로드시에 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성시 유용하다.
	fmt.Println("Main Method start")
}

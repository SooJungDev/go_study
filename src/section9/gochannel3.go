package main

import (
	"fmt"
	"time"
)

func main() {
	// channel
	// 예제1(동기 : 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go :", i)
			time.Sleep(1 * time.Second) // sleep 주석 처리 후 테스트
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}

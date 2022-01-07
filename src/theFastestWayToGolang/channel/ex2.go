package main

import (
	"fmt"
	"time"
)

/**
동기 채널은 보내는쪽에서는 값을 받을때 까지 대기하고, 받는 쪽에서는 채널에 값이 들어올떄까지 대기함
따라서 동기 채널을 활용하면 고루틴의 코드 실행 순서를 제어 할 수 있음
*/
func main() {
	done := make(chan bool) // 동기 채널 생성
	count := 3              // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 고루틴에 true 를 보냄
			fmt.Println("고루틴 :", i)     // 반복문의 변수 출력
			time.Sleep(1 * time.Second) //1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done                   // 채널에 값이 들어올 때가지 대기, 값을 꺼냄
		fmt.Println("메인함수 :", i) // 반복문의 변수 출력
	}
}

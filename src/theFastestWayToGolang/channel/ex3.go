package main

import (
	"fmt"
	"runtime"
)

/**
채널에 버퍼를 1개이상 설정하면 비동기 채널이 생성됨
비동기 채널은 보내는 쪽에서 버퍼가 가득차면 실행을 멈추고 대기하며 받는쪽에서는 버퍼에 값이 없으면 대기함
*/
func main() {
	runtime.GOMAXPROCS(1)
	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4                 // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true            // 채널에 true 를 보냄, 버퍼가 가득차면 대기
			fmt.Println("고루틴 :", i) // 반복문의 변수 출력
		}
	}()

	for i := 0; i < count; i++ {
		<-done                   // 버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인 함수:", i) // 반복문의 변수 출력
	}
}

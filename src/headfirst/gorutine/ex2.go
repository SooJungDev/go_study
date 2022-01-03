package main

import (
	"fmt"
	"time"
)

/**
go 프로그램은 main 고루틴(main 함수를 호출하는 고루틴) 이 종료되면 다른 고루틴이 아직 실행중이더라도
그 즉시 실행을 중단함

따라서 a와 b함수를 실행하는 고루틴이 완료되기 전까지 main 고루틴을 실행중인 상태로 유지해야함
이 작업을 정성개도 하려면 GO의 또다른 기능인 채널을 사용해아함

*/
func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}

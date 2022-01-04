package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
조건 변수는 대기하고 있는 객체 하나만 깨우거나 여러개를 동시에 깨울 때 사용함

Wait() : 고루틴 실행을 멈추고 대기
Signal() : 대기하고 있는 고루틴 하나만 깨움
Broadcast() : 대기하고 있는 모든 고루틴을 깨움
*/
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호시작
			c <- true
			fmt.Println("wait begin :", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end :", n)
			mutex.Unlock() // 무텍스 잠금 해제, con.Wait() 보호 종료
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행 될때까지 기다림
	}

	/*	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금,  cond.Signal() 보호시작
		fmt.Println("signal :", i)
		cond.Signal()
		mutex.Unlock() // 뮤텍스 잠금 해제,  cond.Signal() 보호 종료
	}*/

	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}

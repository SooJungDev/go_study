package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
Once 를 사용하면 함수를 한번 만 실행 할 수 있습니다.
Once는 어떤 상황이든 상관없이 지정된 함수를 딱 한번만 실행시킴
복잡한 반복문 안에서 각종 초기화를 할때 유용함
*/
func hello() {
	fmt.Println("Hello, world")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	once := new(sync.Once) // once 생성

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine :", n)

			once.Do(hello) // 고루틴은 3개지만 hello 함수 한 번만 실행
		}(i)
	}

	fmt.Scanln()
}

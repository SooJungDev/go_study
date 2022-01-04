package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
대기 그룹은 고루틴이 모두 끝날 때까지 기다릴떄 사용합니다. 앞에서는 time.Sleep, fmt.Scanln 함수를 사용하여 고루틴이 끝날 때 까지 임시로 대기함

대기 그룹은 sync.WaitGroup 을 할당한 뒤에 Add,Done,Wait 함수로 사용
고루틴을 생성할때 Add 함수로 고루틴 개수를 추가해줌
고루틴 안에서 Done 함수를 사용하여 고루틴이 끝났다는 것을 알려줌 마지막으로 Wait 함수를 사용하여 모든 고루틴이 끝날떄까지 기다림

Add 함수에는 설정한 값과 Done 함수가 호출되는 횟수는 같아야함
즉 Add(3) 으로 설정했다면 Done 함수는 3번 호출되야함
이 횟수가 맞지 않으면 패닉이 발생하므로 주의해야함

Done 함수는 defer 와 함께 지연 호출로도 사용 할 수 있음
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1)        // 반복할때마다 wg.Add 함수로 1씩 추가
		go func(n int) { // 고루틴 10개 생성
			fmt.Println(n)
			defer wg.Done() // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}
	wg.Wait() // 모든 고루틴이 끝날떄까지 기다림
	fmt.Println("the end")

}

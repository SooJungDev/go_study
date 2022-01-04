package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
원자적 연산은 더이상 쪼갤 수 엇ㅂ는 연산이라는 뜻
따라서 여러 스레드(고루틴) CPU 코어에서 같은 변수(메모리)를 수정할때 서로 영향을 받지 않고 안전하게 연산 할 수 있음
보통 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있음

여러 변수에 고루틴이 동시에 접근하면서 정확하게 연산이 되지 않음
원자적 연산 사용하면 이부분을 해결 할 수 있음
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			// data +=1
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// data -=1
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(data)
}

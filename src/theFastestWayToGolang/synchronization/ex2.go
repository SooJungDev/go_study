package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
읽기, 쓰기 뮤텍스는 읽기 동작과 쓰기 동작을 나누어 잠금(락)을 걸 수 있습니다.

읽기 락(Read Lock): 읽기 락끼리는 서로를 막지 않습니다. 하지만, 읽기 시도 중에 값이 바뀌면 안 되므로 쓰기 락은 막습니다.
쓰기 락(Write Lock): 쓰기 시도중에 다른 곳에서 이전 값을  읽으면 안되고, 다른곳에서 값을 바꾸면 안되므로 읽기,쓰기 락 모두 막습니다.

RLock, RUnlock, Lock, Unlock 함수는 반드시 짝을 맞춰야하며 짝이 맞지 않으면 데드락이 발생하므로 주의
read1,read2 읽기 동작이 모두 끝나야 write 쓰기 동작이 시작된다.
마찬가지로 쓰기 동작이 끝나야 읽기 동작이 시작됨
읽기 동작끼리는 서로를 막지 않으므로 항상 동시에 실행

읽기,쓰기 뮤텍스는 중요한 쓰기 작업을 할 때 다른곳에서 이전 데이터를 읽지 못하도록 방지하거나,읽기 작업을 할때 데이터가 바뀌는 상황을 방지할때 사용함
특히 읽기 쓰기 뮤텍스는 쓰기 동작보다 읽기 동작이 많을떄 유리함

*/
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	var rwMutex = new(sync.RWMutex)

	var data int = 0
	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock() // 쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 1:", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock() // 읽기 뮤텍스 잠금해제, 읽기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 2:", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock() // 읽기 뮤텍스 잠금해제, 읽기 보호 종료
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
채널 이외에도 고루틴의 실행 흐름을 제어하는 동기화 객체를 제공
대표적인 동기화(synchronization) 객체는 다음과 같음

Mutex: 뮤텍스입니다. 상호배제(muual exclusion)라고도 하며 여러 스레드(고루틴) 에서 공우되는 데이터를 보호할때 주로 사용
RWMutex: 읽기/쓰기 뮤텍스입니다. 읽기와 쓰기 동작을 나누어서 잠금(락)을 걸수 있습니다.
Cond: 조건 변수(condition variable)입니다. 대기하고 있느 하나의 객체를 꺠울 수 도 있고 여러개를 동시에 깨울 수 도 있습니다.
Once: 특정 함수를 딱 한번만 실행할대 사용합니다
Pool: 멀티 스레드(고루틴)에서 사용 할 수 있는 객체 풀입니다. 자주 사용하는 객체를 풀에 보관했다가 다시 사용합니다
WaitGroup: 고루틴이 모두 끝날 때까지 기다리는 기능
Atomic: 원자적 연산이라고도 하며 더이상 쪼갤 수 없는 연산이라는 뜻입니다. 멀티스레드(고루틴)멀티코어 환경에서 안전하게 값을 연산하는 기능입니다.

두 고루틴이 경합을 벌이면서 동시에 data 에 접근했기 때문에 append 함수가 정확하게 처리되지 않은 상황
이러한 상황을 경쟁 조건이라고함.

뮤텍스 Lock, UnLock 함수는 반드시 짝을 맞추어주어야 하며 짝이 맞지 않으면 데드락이 발생하므로 주의
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data = []int{} // int 형 슬라이스 생성
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용 할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용 할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))

}

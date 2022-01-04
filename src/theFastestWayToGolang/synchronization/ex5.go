package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { // Data 구조체 정의
	tag    string // 풀태그
	buffer []int  // 데이터 저장용 슬라이스
}

/**
풀은 객체(메모리)를 사용한 후 보관해두었다가 다시 사용하게 해주는 기능
객체를 반복해서 할당하면 메모리 사용량도 늘어나고, 메모리를 해제해야 하는 가비저 컬렉터에게도 부담이 됩니다.
풀은 일종의 캐시라고 할 수 있으며 메모리 할당과 해제 횟수를 줄여 성능을 높이고자 할때 사용합니다.
그리고 풀은 여러 고루틴에서 동시에 사용 할 수 있습니다.

수명주기가 짧은 객체는 풀에 적합하지 않음
*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{ // 풀 할당
		New: func() interface{} {
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)
			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 프ㅜㄹ에 객체 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)

			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()
}

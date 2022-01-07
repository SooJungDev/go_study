package main

import (
	"fmt"
	"time"
)

/**
Go 언어는 여러 채널을  손쉽게 사용 할 수 있도록 select 분기문을 제공
select 분기문은 switch 분기문과 비슷하지만 select 키워드 뒤에 검사 할 변수를 따로 지정하지 않으며
각 채널에 값이 들어오면 해당 case 가 실행됨(close 함수로 채널을 닫았을때도 case가 실행됨)
보통 select 를 계속 처리할 수 있도록 for로 반복해줌 반복하지 않으면 한번만 실행하고 끝냄
switch 분기문과 마찬가지로 select 분기문도 default 케이스를 지정 할 수 잇으며 case 에 지정된 채널에 값이 들어오지 않았을떄 즉시 실행됩니다.
단 default에 적절한 처리를 하지 않으면 cpu 코어를 모두 점유하므로 주의해야함
*/

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}

	}()

	go func() {
		for {
			c2 <- "Hello, world"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1: // 채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
				fmt.Println("c1", i)
			case s := <-c2: // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
				fmt.Println("c2", s)
			case <-time.After(50 * time.Millisecond): //50밀리초 후 현재시간이 담긴 채널이 리턴됨
				fmt.Println("timeout")
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

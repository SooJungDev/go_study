package main

import "fmt"

/**
이전에 채널은 수신 고루틴이 값을 사용하기 전에 송신 고루틴이 값을 보냇음을 보장 할 수 있다고 말한적이 있음
채널은 현재 고루틴의 모든 작업을 중지하는 블록킹으로 이를 보장 할 수 있음
어떤 채널에 대한 송신 연산은 다른 고루틴이 해당 채널에서 값을 가져가기전에 송신 고루틴을 블로킹한다.
고루틴은 자기 자신의 행동을 동기화 할 수 있다. 즉 자신의 실행 타이밍을 조정 할 수 있다.
*/
func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()

}

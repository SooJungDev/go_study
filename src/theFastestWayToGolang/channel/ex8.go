package main

import "fmt"

func num(a, b int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		out <- a   //채널에 a의 값을 보냄
		out <- b   // 채널에 b의 값을 보냄
		close(out) // 채널을 닫음
	}()
	return out // 채널 변수 자체를 리턴
}

func sum3(c <-chan int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		r := 0
		for i := range c { // range 를 사용하여 채널이 닫힐 댸까지 값을 꺼냄
			r = r + i // 꺼낸 값을 모두 더함
		}
		out <- r // 더한 결과를 채널에 보냄
	}()
	return out // 채널 변수 자체를 리턴
}

func main() {
	c := num(1, 2) // 1과 2가 들어 있는 채널이 리턴됨
	out := sum3(c) // 채널 c를 매개변수에 넘겨서 모두 더함
	//더한 값이 들어 있는 out 채널을 리턴
	fmt.Println(<-out) // 3: out 채널에서 값을 꺼냄
}

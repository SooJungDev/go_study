package main

import "fmt"

func sum2(a, b int) <-chan int {
	out := make(chan int) //채널 생성
	go func() {
		out <- a + b
	}()
	return out // 채널 변수 자체를 리턴
}

func main() {
	c := sum2(1, 2)  // 채널을 리턴값으로 받아서 c에서 대입
	fmt.Println(<-c) //3: 채널에 값을 꺼냄
}

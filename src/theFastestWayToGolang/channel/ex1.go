package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b
}

/**
채널을 사용하기 전에는 반드시 make 함수로 공간을 할당해야함 그리고 이렇게 생성하면 동기채널이 생성됨
채널을 매개변수로 받는 함수는 반드시 go 키워드를 사용하여 고루틴으로 실행해야 함
*/
func main() {
	c := make(chan int)

	go sum(1, 2, c) // sum 을 고루틴으로 실행한뒤 채널을 매개변수로 넘겨줌

	n := <-c
	fmt.Println(n)
}

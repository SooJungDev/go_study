package main

import (
	"fmt"
)

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 1; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	// 채널 (channel)

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 데이터 수신(동기) : 채널에서 값 수신 완료 될때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex1 : ", result1)
	fmt.Println("ex1 : ", result2)
	fmt.Println("ex1 : ", result3)
}

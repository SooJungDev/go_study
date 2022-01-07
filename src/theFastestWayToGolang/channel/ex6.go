package main

import "fmt"

/**
보내기 전용 및 받기 전용 채널은 채널 앞 뒤로 <- 연산자를 붙여서 만듬
보통 함수의 매개변수로 사용하거나 구조체의 필드로 사용
보내기 전용(send-only) : chan <- 자료형 형식. c chan <-int는 int형 보내기 전용 채널 c 를 뜻함
보내기 전용 채널은 값을 보낼 수만 있으면 값을 가져오려고 하면 컴파일 에러 발생
받기 전용(receive-only): <-chan 자료형 형식. c <-chan int 는 int형을 받기 전용 채널 c를 뜻함
받기 전용 채널은 range 키워드 또는 <- 채널 형식으로 값으 ㄹ꺼낼 수 만 있으며 값을 보내려고 할때 컴파일 에러가 발생
*/
func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100 // 채널에 값을 보냄
	//fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c) // 채널에 값을 꺼냄
	// c <- 1 채널에 값을 보내면 컴파일 에러

}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}

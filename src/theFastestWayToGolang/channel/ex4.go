package main

import "fmt"

/**
for 반복문 안에서 range 키워드를 사용하면 채널이 닫힐때까지 반복하면서 값을 꺼냅니다.
이미 닫힌 채널에 값을 보내면 패닉이 발생
채널을 닫으면 range 루프가 종료
채널이 열러 있고 , 값이 들어오지 않는다면 range 는 실행되지 않고 계속 대기
만약에 다른 곳에서 채널에 값을 보냈다면(채널의 값이 들어오면) 그때부터 range 가 계속 반복됩니다.
*/
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄
		}
		close(c) // 채널을 닫음
	}()

	for i := range c { // range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i) // 꺼낸 값을 출력
	}
}

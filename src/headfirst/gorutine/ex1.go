package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
고루틴의 실행 시점은 직접 제어할수 없음
매번 다른 순서로 실행됨
일반적인 환경에서 고루틴 간의 전환시기와 전환 기간에 대해 아무것도 보장하지 않습니다.
이는 고루틴이 어떤 방식으로든 가장 호율적으로 실행되도록 만들어줌
하지만 고루틴 간 실행 순서가 중요하다면 채널을 사용하여 고루틴들을 동기화해야함

go 문은 반환 값과 함께 사용 할 수 없음
go 에는 고루틴끼리 서로 통신할 수 있는 방법이 존재하는데 바로 채널(channel) 이다.
채널을 사용하려면 우선 다음 세가지 작업을 할 수 있어야함
- 채널 생성하기
- 채널의 매개변수로 받는 함수 작성하기, 이 함수는 별도의 고루틴에서 실행하여 채널 값을 전달하도록함
- 고루틴에서 채널을 통해 전달한 값 수신하기
*/
func main() {
	go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	time.Sleep(time.Second * 5)
}

func responseSize(url string) {
	fmt.Println("Greeting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(body))
}

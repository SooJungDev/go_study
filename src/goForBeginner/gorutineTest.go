package main

import (
	"fmt"
	"time"
)

/**
메인 함수는 고루틴을 기다려주지 않는다
메인 함수가 실행되는 동안만 실행됨

채널 파이프 같은거 메세지 보낼수도 잇고 받을수 도있음
메세지를 받을때 까지 기다린다.
만약에 2개의 고루틴이 돌아간다면 2개말고 세번째 메세지는 받을수 없음
남아잇는 고루틴이 없는데 하염없이 기다리는 상황이 발생함 (에러남)
매새자 받는거 blocking operation
*/
func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "larry", "soojung"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// go 에러 처리 기초
	// 예러처리 : 소프트웨어의 품질 향상 가장 중요한 것! -> 유형코드 및 에러 정보등을 정보를 남기는 것!
	// Go 에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	// type error interface { Error() string }
	// 즉 Error 메서드를 구현하면 사용자 정의 에러 처리 제작 가능
	// 기본적으로 메서드마다 리턴타입 2개 (리턴값, 에러)
	// 주로 1. Errorrf(에러 타입 리턴) 메소드, 2.Fatal(프로그램 종료) 메소드를 통해서 에러 출력

	// 기본적인 메서드 에러 처리 예제
	f, err := os.Open("unamedfile")
	if err != nil {
		//log.Fatalln(err.Error()) // 방법1
		log.Fatal(err) // 방법2
	}

	fmt.Println("================")
	fmt.Println("ex1 :", f.Name())
}

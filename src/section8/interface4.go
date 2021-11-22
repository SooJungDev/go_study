package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) run() {
	fmt.Println(d.name, " Dog is running!")
}

func (d Cat) run() {
	fmt.Println(d.name, " Cat is running!")
}

func act(animal interface{ run() }) { //익명 인터페이스(타입 정의x)
	animal.run()
}

func main() {
	// 익명 인터페이스 사용예제(즉시 선언 후 사용)

	// 예제1
	dog1 := Dog{"poll", 10}
	cat1 := Cat{"bob", 5}

	// 개 행동실행
	act(dog1)
	// 고양이 행동실행
	act(cat1)
}

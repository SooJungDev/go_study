package main

import "fmt"

func printPointer(myPointer *bool) { // 이 매개변수에 포인터 타입을 사용한다
	fmt.Println(*myPointer) // 전달받은 포인터 주소에 있는 값을 출력
}

func main() {
	var myBool bool = true
	printPointer(&myBool)
}

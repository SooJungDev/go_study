package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한번 선언후 값 변경 금지, 고정된 값 관리 용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d = getHeight() //에러남 상수에서는 함수를 리턴받을 수 없다.
	const e = 35.6
	const f = false
	/**
	  에러 발생
	const g string
	g = "Test3"
	*/

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
}

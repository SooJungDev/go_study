package main

import "fmt"

func main() {
	// 예제1
	a := 30 / 15
	switch a {
	case 2, 4, 6: //a 가 2,4,6 인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: //a가 1,3,5 인 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	// 예제2
	// switch case 자동으로 break 문이 들어가져 있음
	// fallthrough 넣어두면 다음 조건이 실행됨
	switch e := "go"; e {
	case "java":
		fmt.Println("java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "ruby":
		fmt.Println("ruby!")
		fallthrough
	case "c":
		fmt.Println("c!")
	}
}

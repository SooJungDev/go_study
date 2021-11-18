package main

import "fmt"

func main() {
	// 맵 조회 및 순회(iterator)

	// 예제1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println()

	// 예제2 (순서 없으므로 랜덤)
	for k, v := range map1 {
		fmt.Println("ex2:", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2", v)
	}

}

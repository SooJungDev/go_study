package main

import (
	"fmt"
	"os"
)

//에러 체크 방식1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

//에러 체크 방식2
func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의(오류 메세지  확인)
	// 예외 처리 정말 중요
	// 탐색 Seek 중요

	// 파읽 읽기 예제
	// 파일 열기
	file, err := os.Open("/Users/crytal2840/go_study/src/section11/sample.txt")
	errCheck1(err)

	// 읽기 예제1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 회득
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1)", fileInfo, "\n")
	fmt.Println("파일 이름(2)", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3)", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4)", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)
	//fmt.Println(fd1) // 타입 변환 없을 경우
	fmt.Println(string(fd1)) // 타입 변환 한 경우
	fmt.Println("===============================")

	// 읽기 예제2(탐색2 : seek(Offset))
	o1, err := file.Seek(20, 0) // 0: 처음 위치 , 1: 현재위치 2: 마지막 위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n\n")
	fmt.Println("===============================")

	// 읽기 예제3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치로 읽어온다.
	errCheck1(err)

	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n\n")
	fmt.Println("===============================")

	defer file.Close()
}

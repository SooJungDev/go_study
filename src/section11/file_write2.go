package main

import (
	_ "bufio"
	"encoding/csv"
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
	// 파일 쓰기
	// csv 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등의 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일의 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("/Users/crytal2840/go_study/src/section11/test_write.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	//csv writer 생성
	wr := csv.NewWriter(file)
	// wr := csv.newWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"kim", "4.8"})
	wr.Write([]string{"lee", "4.2"})
	wr.Write([]string{"park", "4.4"})
	wr.Write([]string{"cho", "4.45"})
	wr.Write([]string{"hong", "4.9"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 :", fi.Name())
	fmt.Println("운영체제 파일 권한 :", fi.Mode())
}

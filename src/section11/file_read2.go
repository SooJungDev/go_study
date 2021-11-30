package main

import (
	"bufio"
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
	// 파일 읽기
	// csv 파일 읽기 예제
	// 패키지 저장소를 통해서 Excel 등의 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일의 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Open("/Users/crytal2840/go_study/src/section11/sample.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// CSV Reader 생성
	// rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// CSV 내용 읽기(1)
	row1, err1 := rr.Read() //1개의 Row 단위로 읽기
	errCheck1(err1)

	row2, err2 := rr.Read() //1개의 Row 단위로 읽기
	errCheck1(err2)

	fmt.Println("CSV Read Example")
	//fmt.Println(row)
	fmt.Println(row1[0], row1[1], row1[7], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[7], row2[1:5])

	fmt.Println("==============================")

	// CSV 내용 읽기(2)
	rows, err := rr.ReadAll() // 전체 Row 읽기
	errCheck2(err)
	fmt.Println("CSV ReadAll Example")
	//fmt.Println(rows)
	fmt.Println(rows[5][1], " : ", rows[2][1], " : ", rows[6][1])
	fmt.Println("==============================")

	fileInfo, err := file.Stat()
	fmt.Println("파일 정보 출력(1)", fileInfo)
	fmt.Println("파일 이름(2)", fileInfo.Name())
	fmt.Println("파일 크기(3)", fileInfo.Size())
	fmt.Println("파일 수정 시간(4)", fileInfo.ModTime())

	// Row 단위로 csv 파일 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s   ", rows[i][j])
		}
		fmt.Println()
	}
}

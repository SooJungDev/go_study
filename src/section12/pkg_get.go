package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
)

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func rowStuff() {
	filename := "/Users/crytal2840/go_study/src/section12/sample.xlsx"
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet["Sheet1"]
	if !ok {
		panic(errors.New("Sheet not found"))
	}
	fmt.Println("Max row is", sh.MaxRow)
	sh.ForEachRow(rowVisitor)
}

func main() {
	// 사용자 패키지 설치 및 활용 예제
	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1.import 선언 후 폴더 이동후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언
	//  go get github.com/tealeg/xlsx

	// 선언 후 go get 설치 예제(엑셀 파일 읽기)

	fmt.Println("== xlsx package tutorial ==")
	rowStuff()
}

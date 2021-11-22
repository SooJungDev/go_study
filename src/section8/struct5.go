package main

import "fmt"
import "reflect"

type Car struct {
	name string "차량명"
	color string "색상"
	company string "제조사"
}

func main() {
    // 구조체 기본5
	tag := reflect.TypeOf(Car{})

	for i:=0; i<tag.NumField(); i++{
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}

package lib

import "fmt"

func CheckNum(c int32) bool {
	return c > 10
}

func init() {
	fmt.Println("lib Package > Init Method start")
}

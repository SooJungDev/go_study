package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CaculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CaculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제1
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 210000000, 0.025}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	fmt.Println()
	CaculateD(kim)
	CaculateP(&lee)

	fmt.Println("ex1 : ", int(kim.balance))
	fmt.Println("ex1 : ", int(lee.balance))

}

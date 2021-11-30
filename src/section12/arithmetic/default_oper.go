// 두 숫자의 사칙 연산 계산 제공 패키지(1)
package arithmetic

// X, Y 2개의 Integer 구조체
type Numbers struct {
	X int
	Y int
}

// x, y 의 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// x, y 의 차를 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

// x, y 의 곱을 계산해서 반환
func (o *Numbers) Multi() int {
	return o.X * o.Y
}

// x, y 의 나누기를 계산해서 반환
func (o *Numbers) Divide() int {
	return o.X / o.Y
}

package prose

import (
	"fmt"
	"testing"
)

/**
_test.go 중요
go test 명령어는 이 접미사를 가진 이름의 파일을 검색한다

 go test headfirst/automatedTest/prose -v
 모든 테스트 함수의 이름과 테스트 결과가 출력된다.

 -run Two 옵션을 추가하면 이름에 Two 라는 문자열이 포함된 테스트 함수들만 실행됨
 go test headfirst/automatedTest/prose -v -run Two

*/

/**
함수의 이름은 Test 로 시작해야함
Test 뒤로는 아무이름이나 사용 할 수 있음
이 함수에는 testing.T 와 값의 포인터가 전달됨
t.Error -> 테스트 실패를 위해 testing.T 의 메소드를 호출
*/
func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

type testData struct {
	list []string // JoinWithCommas에 전달할 슬라이스
	want string   // 위 슬라이스에 대해 JoinWithCommas 가 반환하길 기대하는 문자ㅕㅇㄹ
}

/**
테이블 주도 테스트
위의 세 테스트 함수를 단위 테스트 함수로 통합할수 있음
테이블 주도 테스트의 가장 큰 장점은 새로운 테스트의 추가가 매우 쉬움
*/
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{}, want: ""},
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) =\"%s\",want \"%s\"", test.list, got, test.want)
		}
	}
}

/**
_test.go 파일에 테스트 함수만 작성 할 수 있는것은 아님
테스트 코드간에 반복되는 코드를 줄이기 위해ㅓ 테스트 파일에서 해당 코드만 따로 "helper" 함수로 분리 할 수 있다.
helper 함수로 따로 분리
문자열 생성코드를 errorString 으로 분리해줌
헬퍼 함수 때문에 코드가 깔금해짐 여전히 동일한 출력값 얻을 수 있음
*/
func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) =\"%s\",want \"%s\"", list, got, want)
}

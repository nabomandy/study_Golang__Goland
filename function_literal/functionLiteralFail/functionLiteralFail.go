package main

import (
	"fmt"
	_ "fmt"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type OpFn func(int, int) int

//func getOperator(op string) func(int, int) int{ // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
func getOperator(op string) OpFn { // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
	if op == "+" {
		return func(a, b int) int { // 람다
			return a + b
		} // add라는 애의 주소 위에 add 함수
	} else if op == "*" {
		return func(a, b int) int { // 람다
			return a * b
		}
	} else { // + 나 * 가 아니면 nul 반환
		return nil
	}

}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입
	//var operator func(int, int) int // operator라는 이름의 변수, 타입은 func(int, int) int
	var operator OpFn // operator라는 이름의 변수, 타입은 func(int, int) int
	operator = getOperator(+) // 컴파일 에러 : <expression> expected, got ')'
	operator = getOperator('+') // 컴파일 에러 : Cannot use ''+'' (type rune) as the type string

	result := operator(3, '4') // 컴파일 에러x but, 55 -> 3 + 52(아스키 코드 번호)
	result := operator(3, "4") // 컴파일 에러 : Cannot use '"4"' (type string) as the type int
	fmt.Println(result)

}

/**
함수리터럴(람다식)-> 함수를 박아 넣었다

var a int = 3
b := a
를 바꾸면
b := 3


func add(a, b int) int {
	return a + b
}
f := add
를 바꾸면
f := (a, b int) {
	return a + b
}

함수 리터럴은 일반 함수랑 다르다
-> 함수 리터럴은 내부 상태(First class, First Object)를 가질 수 있다. 일반함수는 상태를 가질 수 없다.





*/

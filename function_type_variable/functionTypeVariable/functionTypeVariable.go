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

//func getOperator(op string) func(int, int) int{ // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입
//-> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
func getOperator(op string) OpFn { // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입
	// -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
	if op == "+" {
		return add // add라는 애의 주소 위에 add 함수
	} else if op == "*" {
		return mul
	} else { // + 나 * 가 아니면 nul 반환
		return nil
	}

}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입
	//var operator func(int, int) int // operator라는 이름의 변수, 타입은 func(int, int) int
	var operator OpFn // operator라는 이름의 변수, 타입은 func(int, int) int
	operator = getOperator("+")

	result := operator(3, 4) // = add(3, 4)와 똑같다
	fmt.Println(result)

}

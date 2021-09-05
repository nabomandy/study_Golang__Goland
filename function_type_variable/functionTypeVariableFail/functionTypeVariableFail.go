package main

//
//import (
//	"fmt"
//	_ "fmt"
//)
//
//func add(a, b int) int {
//	return a + b
//}
//
//func mul(a, b int) int {
//	return a * b
//}
//
//type OpFn func(int, int) String // type OpFn func(int, int) int 가 되어야 한다. -> 람다 변환 간 실수로 작성한 오류
///**
//런타임 오류 : 반환 타입이 int 인데 String으로 잘못한 경우. 먼저 컴파일 오류가 난다.
//# command-line-arguments
//.\functionTypeVariableFail.go:16:26: undefined: String
//.\functionTypeVariableFail.go:21:3: cannot use add (type func(int, int) int) as type OpFn in return argument
//.\functionTypeVariableFail.go:23:3: cannot use mul (type func(int, int) int) as type OpFn in return argument
//*/
//
////func getOperator(op string) func(int, int) int{ // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
//func getOperator(op string) OpFn { // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
//	if op == "+" {
//		return add // add라는 애의 주소 위에 add 함수
//	} else if op == "*" {
//		return mul
//	} else { // + 나 * 가 아니면 nul 반환
//		return nil
//	}
//
//}
//
//func main() {
//	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입
//	//var operator func(int, int) int // operator라는 이름의 변수, 타입은 func(int, int) int
//	var operator OpFn // operator라는 이름의 변수, 타입은 func(int, int) int
//	operator = getOperator("+")
//
//	result := operator(3, 4) // = add(3, 4)와 똑같다
//	fmt.Println(result)
//
//}

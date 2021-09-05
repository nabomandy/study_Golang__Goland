package main

import (
	"fmt"
	_ "fmt"
)

func add(a, b int) int {
	return a + 'b' // 컴파일 에러, 런타임 에러는 나지 않는다. 단지 'b'는 아스키로 변환되어 예상한 값과 다르게 나온다. 결과값 : 101
}

func mul(a, b int) int {
	return a * b
}

type OpFn func(int, int) int

/**
func getOperator(op string) func(int, int) int{ // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
*/
func getOperator(op string) OpFn { // op에 따른 함수 타입 반환 -> func(int, int) int 전체가 하나의 타입 -> 이렇게만 쓰면 보통 지저분해서 별칭 타입을 만들어준다.
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

	operator = getOperator("s") // 컴파일 에러, 런타임 에러가 나지 않는다. " "안에 "+" 또는 "*"가 들어가야 하는데, "s"도 스트링으로 들어간 건 마찬가지이기 때문이다.

	/**
	런타임 에러 :
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x0 pc=0x49c516]

	goroutine 1 [running]:
	main.main()
		C:/Users/andya/GolandProjects/awesomeProject/function_type_variable/functionTypeVariableFail/functionTypeVariableFail2.go:36 +0x16

	*/

	result := operator(3, 4) // = add(3, 4)와 똑같다
	fmt.Println(result)
}

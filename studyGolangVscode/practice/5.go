// 멀티반환함수

package main

import "fmt"

func Divide(a, b int) (int, bool) { // 함수 선언
	if b == 0 {
		return 0, false // 제수가 0일 때 반환
	}
	return a / b, true // 제수가 0이 아닐 때 반환
}

func main() {
	c, success := Divide(9, 3) // 제수가 0이 아닌 경우
	fmt.Println(c, success)
	d, success := Divide(9, 0) // 제수가 0인 경우
	fmt.Println(d, success)
}

/*
이 함수는 int 타입 a,b를 매개변수로 받고 int 타입과 bool 타입을 반환한다.
(a int, b int)와 같이 매개변수 타입이 같으면 간단히 (a, b int)처럼 표현할 수 있다.


결과값
3 true
0 false
*/

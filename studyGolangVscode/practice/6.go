// 멀티반환함수 반환할 변수명 명시

package main

func Divide(a, b int) (result int, success bool) { // 함수선언, 반환할 변수명 명시
	if b == 0 {
		return 0, false // 제수가 0일 때 반환
	}
	return a / b, true // 제수가 0이 아닐 때 반환
}

// ** 반환할 변수에 이름을 지정할 경우, 모든 반환 변수에 이름을 지정해야 한다. 모두 지정하거나, 모두 지정하지 않거나!

/*
이 함수는 int 타입 a,b를 매개변수로 받고 int 타입과 bool 타입을 반환한다.
(a int, b int)와 같이 매개변수 타입이 같으면 간단히 (a, b int)처럼 표현할 수 있다.

(result int, success bool) { // 반환할 변수명 명시

결과값
3 true
0 false
*/

package main

import "fmt"

func main() {
	var a int = 10 // a 변수 선언.
	// var           +   a    + int  = 10
	// 변수선언키워드 + 변수명 + 타입 = 초깃값
	var msg string = "Hello Variable" // msg 변수 선언
	// 변수를 사용하려면 변수를 선언해야 한다. 변수 선언은 메모리 할당이다.

	a = 20               // a 값 변경
	msg = "Good Morning" // msg 값 변경
	fmt.Println(msg, a)  // msg와 a 값 출력
}

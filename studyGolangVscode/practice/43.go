package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)
}

/*
결과값
p1 == p2: true
p2 == p3: false
*/

/*
포인트 연산에서 equal == 을 쓸 수 있다.
두 포인터가 가지고 있는 값이 같은지를 물어보는 거고
두 포인터가 같은 메모리를 가리키고 있는지
*/

// 포인터 이해를 위한 예제

package main

import (
	"fmt"
)

func main() {
	var a int = 500
	var p *int

	p = &a // -> p는 a의 주소값

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // *p -> 포인터 앞에 *을 붙이면 그 공간. 여기서는 우변으로 적용되었기 떄문에 값이 된다.

	*p = 100 // 포인터 p가 가지고 있는 값이 나타내는 공간을 100으로 바꿔라 -> a의 값을 100으로 바꿔라
	fmt.Printf("a의 값: %d\n", a)
}

/*
결과값
p의 값: 0xc00010a000 -> a의 메모리 주소 값
p가 가리키는 메모리의 값: 500
a의 값: 100

*/

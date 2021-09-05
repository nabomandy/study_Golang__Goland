package main

import (
	"fmt"
	_ "fmt"
)

func aaa() int {

	i := 2
	return i + 10 // 컴파일 에러 : Unresolved reference 'i' -> int 타입으로 리턴 값을 주게 이용
}

func main() {
	i := 0

	f := func() { // function인데 입력값이 없다, 이 함수 리터를은 외부에서 변수 상태를 가져올 수 있다. 캡쳐할 수 있다.
		i += 10
	}

	i++ // 컴파일 에러 : '++' unexpected -> ++i에서 i++로 변경

	f()

	fmt.Println(i)

}

/**
이 함수 리터를은 외부에서 변수 상태를 가져올 수 있다. 캡쳐할 수 있다.
캡쳐는 값 복사가 아닌 레퍼런스 복사이다. 고에서는 레퍼런스가 없다.
즉, 포인터가 복사된다고 볼 수 있다. 같은 메모리 공간을 가리킨다.

*/

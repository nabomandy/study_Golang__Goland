package main

import (
	"fmt"
	_ "fmt"
)

func CaptureLoop() {
	f := make([]func(), 3) // []func() 입력도 없고 출력도 없는 func 타입 의 슬라이스 3개 만들어요
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) //캠쳐된 i값 출력
		}
	}
	// 이 시점에서 i = 3
	for j := 0; j < 3; j++ {
		f[j]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3) //리터럴 3개를 갖고 있는 슬라이스
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}

/**
이 함수 리터를은 외부에서 변수 상태를 가져올 수 있다. 캡쳐할 수 있다.
캡쳐는 값 복사가 아닌 레퍼런스 복사이다. 고에서는 레퍼런스가 없다.
즉, 포인터가 복사된다고 볼 수 있다. 같은 메모리 공간을 가리킨다.

이 부분을 잘 모르면 고루틴에서 문제가 많이 된다.
고루틴 쓸 때 람다로 다른 루틴으로 값을 넘길 수 있는데(캠쳐로 넘긴다)
그래서 값을 복사하는게 아니라 레퍼런스를 복사하는구나를 잘 알고 있어야 한다.
*/

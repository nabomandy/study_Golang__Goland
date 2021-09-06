package main

import (
	"fmt"
	"unsafe" // 표준패키지. 안전하지 않은 함수들을 제공
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	var a int = 45
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a)) //Sizeof 얼마나 할당되었느냐 사이즈를 제공
}

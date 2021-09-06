package main

import (
	"fmt"
	"unsafe" // 표준패키지. 안전하지 않은 함수들을 제공
)

type User struct {
	A int8 // 1바이트
	B int  //8 바이트
	C int8 // 1바이트
	D int  //8 바이트
	E int8 // 1바이트
} // 19바이트

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

// 결과값40
// 19바이트인데 40바이트를 쓴다 -> 메모리 낭비

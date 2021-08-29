package main

import (
	"fmt"
	_ "fmt"
)

type myInt int // 사용자 정의 별칭 타입

// myInt 별칭타입을 리시버로 갖는 메서드
func (a myInt) add(b int) int {
	//return a + b //-> why is it er?
	// .\main.go:12:11: invalid operation: a + b (mismatched types myInt and int)
	// 에러메시지를 보면 음 이게 문제구나 알겠는데, 왜지?  myInt는 int 타입 아닌거? 엄연히 다른 타입으로 Go는 생각하는거면 뭐라고 생각하는 걸까?
	return int(a) + b
}

func main() {
	var a myInt = 10       // myInt 타입 변수
	fmt.Println(a.add(30)) //myInt 타입의 add() 메서드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50)) // int 타입을 타입 변환
}

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User { // *User 로 반환한다
	var u = User{name, age}
	return &u // u의 주소를 반환한다
}

func main() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}

/*
u는 { } 안에서만 존재한다. 그래서 C나 C++에서는 위와 같은 코드의 경우 에러가 난다. ->Dangling
func 내에서 있는 지역변수는 스택 메모리에 쌓이게 된다. func 이 돌아갈 때 팝을 해버린다 -> 끝날 때 없애버린다. -> 무효한 주소가 된다.

Go에서는 Escape Analyzing(탈출분석)을 통해서 스택에 만들지 힙에 만들지를 결정한다.
힙에 만든다. -> 힙은 프로그램 전체에서 사용되는 영역. -> 쓰임이 다하면 사라진다. -> 쓰임이 있는 한 사라지지 않는다.

*/

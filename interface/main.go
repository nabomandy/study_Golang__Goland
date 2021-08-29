package main

import (
	"fmt"
	_ "fmt"
)

/***
Go 언어에서는 -er을 붙여서 인터페이스명을 만드는 것을 권장하고 있습니다. String() 메서드를 가진 인터페이스란 뜻으로 Stringer 라고 만들었다.
Stringer의 원뜻과는 관계가 없다.
*/

type Stringer interface { // Stringer 인터페이스 선언
	String() string // 매개변수 없이 string 타입을 반환하는 String()메서드를 포함
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // Student의 String() 메서드

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // 문자열 만들기 Sprintf 와 printf 의 차이 확인해볼 것
}

func main() {
	student := Student{"앤디", 31} // Student 타입
	var stringer Stringer        // Stringer 타입

	stringer = student // stringer 값으로 student 대입

	fmt.Printf("%s\n", stringer.String()) // stringer 의 String() 메서드 호출

}

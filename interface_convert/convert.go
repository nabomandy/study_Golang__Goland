package main

import (
	"fmt"
	_ "fmt"
)

type Stringer interface { // 인터페이스
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student) //Student 타입으로 타입변환
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}

	PrintAge(s)

}

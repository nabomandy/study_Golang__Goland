package main

import (
	"fmt"
	_ "fmt"
)

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read()")
}

//func (f *File) Close() {
//	fmt.Println("Close()")
//}

func ReadFile(reader Reader) {
	c, ok := reader.(Closer) // reader 인터페이스를 Closer로 타입변환 -> 서로 다른 인터페이스로 변환이 가능하다. 왜지?
	fmt.Println(c, ok)
	if ok {
		c.Close()
	}

}

func main() {
	file := &File{}
	ReadFile(file)
}

/**
인터페이스는 관계를 나타낸다. 구현이 빠진 관계. 메서드의 목록을 가지고 있다.
이 메서드를 가지고 있는 객체면 전부 다 이 인터페이스로 사용이 가능하다.
고는 덕타이핑을 지원한다.


*/

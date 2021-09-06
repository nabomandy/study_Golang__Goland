package main

import (
	"container/list"
	_ "container/list"
	"fmt"
	_ "fmt"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4) // 4라는 요소가 삽입된다.
	e1 := v.PushBack(1)
	v.InsertBefore(3, e4) // e4 전에 넣어라
	v.InsertAfter(2, e1)  // e1 다음에 넣어라

	for e := v.Front(); e != nil; e = e.Next() { // e := v.Front() 는 맨앞에 요소를 의미
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { // e := v.Back()은 맨 뒤에 값이 나온다
		fmt.Print(e.Value, " ")

	}

}

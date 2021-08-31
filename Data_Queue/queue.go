package main

import (
	"container/list"
	_ "container/list"
	"fmt"
	_ "fmt"
)

type Queue struct { // Queue 구조체 정의
	v *list.List
}

func (q *Queue) Push(val interface{}) { //요소 추가
	q.v.PushBack(val) // 값을 맨 뒤에 넣는다.
}

func (q *Queue) Pop() interface{} { // 요소를 반환하면서 삭제
	front := q.v.Front() // 맨 앞에 있는 애
	if front != nil {    // 빈 리스트가 아니면
		return q.v.Remove(front) //리무브 하면 리무브된 엘리먼트의 밸류가 리턴된다.
	}
	return nil // 비어있으면 닐을 반환한다.
}

func NewQueue() *Queue { // 새로운 큐를 만드는 거
	return &Queue{list.New()} // Queue 객체에 인스턴스를 반환
}

func main() {
	queue := NewQueue() // 새로운 큐 생성
	//queue := NewQueue{} // 새로운 큐 생성

	for i := 1; i < 5; i++ { // 요소 입력 1부터 5개를 넣어볼게요
		queue.Push(i)
	}
	v := queue.Pop() // 뽑아서
	for v != nil {
		fmt.Printf("%v -> ", v) // 무슨 타입인지는 모르니까 빈 인터페이스라서
		v = queue.Pop()
	}

}

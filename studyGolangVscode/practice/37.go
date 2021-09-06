package main

import "fmt"

type Student struct {
	Age   int // 대문자로 시작하는 필드는 외부로 공개된다.
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이:%d 번호:%d 점수:%.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2}

	// student 구조체 모든 필드가 student2로 복사된다.
	student2 := student

	PrintStudent(student2) // 함수 호출 시에도 구조체가 복사된다.

}

/*
Go 언어 내부에서는 필드 각각이 아닌 구조체 전체를 한 번에 복사한다.
대입 연산자가 우변 값을 좌변 메모리 공간에 복사할 때 ' 복사되는 크기'는 '타입 크기'와 같다. 구조체 크기는 모든 필드를 포함하므로 구조체 전체 필드가 복사되는 거다.
*/

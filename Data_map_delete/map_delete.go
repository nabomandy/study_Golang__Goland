package main

import "fmt"

func main() {
	// 맵 생성
	m := make(map[int]int)

	// 요소 추가
	m[1] = 0
	m[2] = 2
	m[3] = 3

	// 요소 삭제 - delete
	delete(m, 3)      // 요소 삭제
	delete(m, 4)      // 없는 요소 삭제 시도 -> 에러가 나지는 않는다.
	fmt.Println(m[3]) // 삭제된 요소 값 출력 -> 기본 값인 0이 나온다.
	fmt.Println(m[2]) // 존재하는 요소값 출력
	fmt.Println(m[4]) // 없는 요소 출력 시도 -> 0이 나온다.

	// 값 존재여부 확인 은 아래 코드인데, 어떻게 쓰는 걸까.
	v, ok := m[3]
	fmt.Println(v, ok) // 이게 없으면 v, ok에 밑줄
}

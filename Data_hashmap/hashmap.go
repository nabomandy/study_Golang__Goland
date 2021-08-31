// 해시함수로 맵을 만들어 보자
package main

import (
	"fmt"
	_ "fmt"
)

const M = 10 // 나머지 연산의 분모

func hash(d int) int {
	return d % M // 나머지 연산
}

func main() {
	m := [M]int{}     // 값을 저장할 배열
	m[hash(23)] = 10  // 키 23에 값 설정
	m[hash(259)] = 50 // 키 259에 값 설정

	// 값 출력
	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%v = %v\n", 259, m[hash(259)])

}

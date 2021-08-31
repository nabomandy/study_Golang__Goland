// 정렬되지 않은 값이 출력
// 맵은 내부에서 보관할 때 입력한 순서와도, 키 값과도 상관 없이 데이터를 보관한다. -> 출력 순서와 입력한 순서가 같지 않다는 의미이다.

package main

import (
	"fmt"
	_ "fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product) // 맵 생성

	// 요소 추가
	m[16] = Product{"볼펜", 500}
	m[48] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}
	m[897] = Product{"키중복샤프심", 100000} // 더 아래에 있는 중복 키값이 출력됨. 에러는 안남
	m[48] = Product{"중복자", 3000}
	m[123] = Product{"", 3000}

	for k, v := range m {
		fmt.Println(k, v)
	}

}

/*
결과 값
16 {볼펜 500}
48 {중복자 3000}
78 {자 1000}
345 {샤프 3000}
897 {샤프심 500}
123 { 3000}
*/

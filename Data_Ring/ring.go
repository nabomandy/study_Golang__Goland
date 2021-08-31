package main

import (
	_ "container/list"
	"container/ring"
	_ "fmt"
)

func main() {
	r := ring.New(5) // 요소가 5개인 링을 생성

	n := r.Len() // 링 길이 반환

	for i := 0; i < n; i++ {
		r.Value = 'A' + i // 순화하면 모든 요소에 값 대입
		r = r.Next()
	}

}

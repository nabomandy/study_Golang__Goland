// 서브 고루틴이 종료될 때까지 기다리기
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {

	for i := a; i <= b; i++ {
		sum = a + b
		fmt.Println(sum)
		fmt.Println("p")
	}
}

func main() {
	SumAtoB(3, 4)
}

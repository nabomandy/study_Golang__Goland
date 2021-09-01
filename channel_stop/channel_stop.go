package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	ch := make(chan int) // 크기 0인 채널 생성tit

	ch <- 9                    // 9를 넣는다. main() 함수가 여기서 멈춘다
	fmt.Println("Never print") // 실행되지 않는다.
	ints := ch()
	fmt.Println(ints)

}

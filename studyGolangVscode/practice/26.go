//무한루프
// 1초마다 하나씩 숫자가 찍힌다. 무한히.

package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for {
		time.Sleep(time.Second) // Sleep은 프로그램을 잠시 멈춰라 -> 1초 동안 멈춰라 -> 정확하게는 쓰레드를 멈춰라 라는 의미
		fmt.Println(i)
		i++
	}
}

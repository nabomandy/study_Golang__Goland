package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	//hanguls := []rune{ '가', '나', '다', 라, '마', '바', '사' }
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumber() {
	for i := 1; i < 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul() // 새로운 고루틴 생성
	go PrintNumber() // 새로운 고루틴 생성

	time.Sleep(3 * time.Second) // 3초간 대기

}

package main

import (
	"fmt"
	_ "fmt"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

func main() {
	var wg sync.WaitGroup // 서브 고루틴이 종료될 때까지 대기하는 방법
	ch := make(chan int)  // 채널생성

	wg.Add(1)          // 작업개수 1로 설정
	go square(&wg, ch) // 고루틴 생성
	ch <- 9            // 채널에 데이터 넣음
	wg.Wait()          // 작업이 완료되길 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 데이터 빼옴

	time.Sleep(time.Second) // 1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done() // 모든 작업이 완료될 떄까지 대기했다가 프로그램을 종료
}

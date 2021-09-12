package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile(filename)
	if err != nil {
		return err // 혹은 다른 에러 처리
	}

	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		// 읽은 num 값 사용

	}

}

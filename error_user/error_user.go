// 사용자 에러 반환 방법 두가지
//fmt.Errof -> 반환 타입이 error
//errors.New(text string) error 함수 사용. 단 얘는 인자를 못 넣는다.

package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f) // f가 음수이면 에러 반환
		//return 0, error.New("제곱근응 양수여야 합니다") // 이것도 가능
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // 에러 출력
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}

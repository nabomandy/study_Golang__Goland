package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	//hanguls := []rune{ '가', '나', '다', 라, '마', '바', '사' } // 컴파일 에러 : Unresolved reference '라'
	/**
	런타임 에러 :
		./prog.go:10:42: undefined: 라
	*/

	for _, v := range hanguls {
	//for a, v := range hanguls { // 컴파일 에러 : Unused variable 'a -> a가 사용되는 곳이 없다. -> a가 필요 없으니 _로 하면 된다.
	/**
	런타임 에러 :
		./prog.go:16:6: i declared but not used
	 */
		time.Sleep(300 * time.Millisecond)
		//time.Sleep("300" * time.Millisecond) // 컴파일 오류 : Invalid operation: "300" * time.Millisecond (mismatched types string and Duration)
		fmt.Printf("%c", v)
		/**
		런타임 에러 :
		./prog.go:23:20: invalid operation: "300" * time.Millisecond (mismatched types untyped string and time.Duration)

		Go build failed.
		 */
	}
}

func PrintNumber() {
	//for i := 1; i < 5; i++ {
	for i := 1 i < 5; i++ { // ;가 빠졌다.
		/**
		컴파일 에러 :
		'!=', '%', '&', '&^', '(', '*', '+', '-', '/', <, <<, <=, '==', '>',
		'>=', '>>', '[', '^', '{', '|' or '~' expected, got 'i'
		Unresolved reference 'i'
		i < 5 evaluated but not used

		 */
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul() // 새로운 고루틴 생성
	go PrintNumber() // 새로운 고루틴 생성

	time.Sleep(3 * time.Second) // 3초간 대기
}

/**
논술의 기본은 생각을 하는 힘을 기르는 것입니다. 생각을 하는 힘은 차이를 인정하는 것부터 시작합니다. 다름을 인정하게 되면 남들과 다른 자신의 생각이 이상한 것이 아님을 깨닫게 됩니다.
궁극적으로 수업에 참가하는 청소년들이 행복해지길 바랍니다.

 */

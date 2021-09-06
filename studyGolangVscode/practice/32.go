// 이중 조건문 전체를 빠져나가기 위해서 레이블을 활용하는 방법

package main

import "fmt"

func main() {
	a := 1 // 선언, 초기화
	b := 1

OuterFor: // 어떠한 이름을 적고 :만 딱 적으면, 얘는 레이블이다
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // OuterFor까지 브레이크 한다.
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

}

package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Print("김일등님 평균점수는", (math+eng+history)/3, "입니다.")

	math := 88
	eng := 92
	history = 53
	fmt.Print("송이등님 평균점수는", (math+eng+history)/3, "입니다.")

	math = 78
	eng = 73
	history = 78
	fmt.Print("박상동님 평균 점수는", (math+eng+history)/3, "입니다.")
}

// 중복코드가 발생한 bad coding.

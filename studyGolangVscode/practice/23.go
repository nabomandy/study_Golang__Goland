package main

import "fmt"

type ColorType int // 별칭 ColorType을 선언하고 const 열거값 정의
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// int로 써도 되지만 굳이 ColorType로 타입을 새로 해서 하는 이유는 -> 용도를 명확하게 하기 위해서
// 컬러를 나타내기 위한 코드 값으로 쓰겠다는 의미

// 각 ColorType 열거값에 따른 문자열을 반환하는 함수
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}

/*

결과값
My favorite color is Blue

*/

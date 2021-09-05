package main

import "fmt"

var gvar int //not an error

func main() {
	var one int   // 컴파일 에러 : unused variable
	two := 2      //컴파일 에러 : unused variable
	var three int // 컴파일 에러 : even though it's assigned 3 on the next line
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}

/**
런타임 에러 :
# awesomeProject/mission/UnusedVar/unusedVar
.\unusedVar.go:8:6: one declared but not used
.\unusedVar.go:9:2: two declared but not used
.\unusedVar.go:10:6: three declared but not used

*/

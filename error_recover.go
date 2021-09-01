package main

import "fmt"

m
import (
	"fmt"
	_ "fmt"
)

func f()  {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil { //인자가 없고 출력도 없는 함수 리터럴
			fmt.Println("panic 복구 -", r)
		}
	}()
	
	g()
	fmt.Println("f() 함수 끝")
	
}


func g() {
	fmt.Printf()
}


func main() {

}

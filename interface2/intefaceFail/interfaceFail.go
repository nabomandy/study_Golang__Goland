package main

import _ "testing"

type Stringer interface {
	String() string    // 컴파일 에러 : Duplicate method 'String'
	String(int) string // String 이라는 메서드명이 겹친다. 컴파일 에러 : Duplicate method 'String'
	//StringFail(int) string { // 인터페이스에서는 메서드 구현을 포함하지 않는다.
	//	return a+b
}

/**
컴파일 에러 :
# command-line-arguments
.\interfaceFail.go:8:25: syntax error: unexpected {, expecting semicolon or newline or }
.\interfaceFail.go:12:1: syntax error: non-declaration statement outside function body
*/

func main() {

}

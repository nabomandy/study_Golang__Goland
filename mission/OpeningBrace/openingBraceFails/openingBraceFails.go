package main


import "fmt" // 컴파일 에러 : Unused import

func main()
{ // 컴파일 에러 : Missing function body. error, can't have the opening brace on a separate line
	fmt.Println("Hello there!")

}

/**
Compile Error:
# command-line-arguments
.\openingBraceFails.go:5:6: missing function body
.\openingBraceFails.go:6:1: syntax error: unexpected semicolon or newline before {
*/

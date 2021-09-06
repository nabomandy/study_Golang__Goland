package main

import "fmt"

func main() {
	str := "Hello 월드!" // 한영이 섞인 문자열
	arr := []rune(str) // 문자열을 []rune으로 타입 변환

	for i := 0; i < len(arr); i++ { // 문자열 크기를 얻어 순회
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])

	}
}

/*

타입:int32 값:72 문자값:H
 타입:int32 값:101 문자값:e
 타입:int32 값:108 문자값:l
 타입:int32 값:108 문자값:l
 타입:int32 값:111 문자값:o
 타입:int32 값:32 문자값:
 타입:int32 값:50900 문자값:월
 타입:int32 값:46300 문자값:드
 타입:int32 값:33 문자값:!

*/

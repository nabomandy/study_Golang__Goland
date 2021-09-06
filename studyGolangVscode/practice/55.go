package main

import "fmt"

func main() {
	str := "Hello 월드!" // 한영이 섞인 문자열

	for i := 0; i < len(str); i++ { // 문자열 크기를 얻어 순회

		// 바이트 단위로 출력
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])

	}
}

/*

결과값
 타입:uint8 값:72 문자값:H
 타입:uint8 값:101 문자값:e
 타입:uint8 값:108 문자값:l
 타입:uint8 값:108 문자값:l
 타입:uint8 값:111 문자값:o
 타입:uint8 값:32 문자값:
 타입:uint8 값:236 문자값:ì
 타입:uint8 값:155 문자값:
 타입:uint8 값:148 문자값:
 타입:uint8 값:235 문자값:ë
 타입:uint8 값:147 문자값:
 타입:uint8 값:156 문자값:
 타입:uint8 값:33 문자값:!


*/

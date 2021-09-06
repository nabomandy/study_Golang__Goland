package main

import (
	"bufio"
	_ "fmt"
	"os"
	_ "os"
)

func Readfile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err // -> 위에서 반환이 스트링하고 에러니까.
	}

	defer file.Close()             // -> file은 한번 열리면 file 핸들을 닫아줘야 한다.
	rd := bufio.NewReader(file)    // 리더 인스턴스로 데이터를 읽어오는거
	line, _ := rd.ReadString('\n') // \n 한줄 띄우는 거 까지 읽는 거니까 한줄에 읽겠다는 의미. ReadLine과 같은 기능이긴 한데 ReadLine 같은 경우에는 반환타입이 byte 슬라이스에 에러에 그래서 그냥 String으로 쓴다.
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
}

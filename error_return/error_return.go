package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	//file, err := os.OpenFile(filename)
	file, err := os.Open(filename)
	if err != nil { // 에러가 nil 이 아니면-> 에러가 있ㅇ면
		return "", err // 에러나면 에러반환 -> 위에 반환 타입이 스트링 에러니까.
	}
	defer file.Close()
	rd := bufio.NewReader(file)    // 파일 내용 읽기, 뉴리더는 인자로 io리더를 받는다 인터페이스를 받는다.
	line, _ := rd.ReadString('\n') // 한줄 띄는 거 까지 읽는다 -> 한줄을 읽는다 -> readline하고 같긴한데 반환이 바이트슬라이스라서 패쓰
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) // 파일 생성 -> 이 때 에러는 대게 권한이 없거나 디스크 용량이 없거나 그러겠죠?
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) // 파일에 문자열 쓰기 // Fprintln은 첫번째 인자로 어디에 쓸 건지 정할 수 있다.
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) // 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") // 파일 생성
		if err != nil {                                // 에러를 처리
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename) // 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line)
}

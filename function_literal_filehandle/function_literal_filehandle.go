package main

import (
	"fmt"
	_ "fmt"
	"os"
	_ "os"
)

type Writer func(string)
type WriterInterface interface {
	Write(string)
}

func writeHello(writer Writer) { // writer 함수 타입 변수 호출
	writer("Hello World")
}

func writeHello2(writer WriterInterface) {
	writer.Write("Hello World")

}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	//writeHello(func(msg string) {
	//	fmt.Fprintln(f, msg) // 함수 리터럴 외부 변수 f 사용 di 주입
	//})

	writeHello(func(msg string) {
		fmt.Println(msg)
	})

}

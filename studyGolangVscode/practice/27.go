// 무한 for 문

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // stdin이라는 변수를 만드는데, 얘는 bufio의 뉴 리더의 결과이다.
	for {
		fmt.Println("숫자를 입력하세요")
		var number int
		_, err := fmt.Scanln(&number) // _는 읽어온 값의 갯수 err는 에러여부. _를 쓴 것은 빈칸지시자라고 해서 첫번째 리턴값을 안쓰겠다는 것. _을 안쓰면 에러가 난다.
		if err != nil {
			fmt.Println("숫자로 입력하세요") // -> ""는 문자열

			// 카보드 버퍼를 지웁니다.
			stdin.ReadString('\n') // -> '' 는 문자 하나
			continue
		}
		fmt.Println("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 { // 2로 나눈 나머지가 0이면 -> 짝수이면 -> 짝수 검사를 합니다.
			break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}

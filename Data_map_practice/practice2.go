package main

import (
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "os"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	m := make(map[int]Student)

	m[5] = Student{"최번개", 67}
	m[17] = Student{"송하나", 89}
	m[23] = Student{"화랑", 97}
	m[38] = Student{"백두산", 78}
	m[45] = Student{"김갑환", 56}

	fmt.Println("38번:", m[38])

}

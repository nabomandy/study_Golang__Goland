package main

import "fmt"

// 상수값에 코드를 부여.
const Pig int = 0
const Cow int = 1
const Chicken int = 2

// 코드값에 따라서 다른 텍스트를 출력
func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음매")
	} else if animal == Chicken {
		fmt.Println("꼬기오")
	} else {
		fmt.Println("... ")
	}
}

// PrintAnimal() 함수를 호출하여 동물 소리를 출력
func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(7)

}

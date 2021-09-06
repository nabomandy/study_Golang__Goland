package main

import (
	"bufio"
	"calculator/calc"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		calculator      *calc.Calculator   = calc.NewCalculator() // 계산기 struct
		operationUnit   calc.OperationUnit                        // OperationUnit이라는 interface type을 통해 Polymorphism 이용
		operationResult float64                                   // 계산 결과를 담음
		operationErr    error                                     // 계산 수행에 대한 error을 담음
	)
	// main package에서는 calc package에 정의된 unexported name인 id에 접근할 수 없다.
	//calculator.id = "Jinsu Park"

	scanner := bufio.NewScanner(os.Stdin)

	// 연산 option을 위해 값을 입력받기 (e.g. 1)
	fmt.Printf(`A simple calculator program.
=======================================================
Operation options

    1. Mulitply a, b float 64
    2. Sqaure val, square float64

Please input an int for your desired operation.
>>> `)
	scanner.Scan()
	option, _ := strconv.Atoi(scanner.Text())

	// args를 위해 입력받기 (e.g. 10 20)
	fmt.Printf("Please input floats as args.\n>>> ")
	scanner.Scan()
	inputs := make([]float64, 0)
	for _, input := range strings.Split(scanner.Text(), " ") {
		f, _ := strconv.ParseFloat(input, 64)
		inputs = append(inputs, f)
	}

	// 연산에 대한 multiplexing. 즉 option에 따른 연산을 수행한다는 의미
	// operationUnit이라는 OperationUnit interface type을 통해 Polymorphism 이용
	switch option {
	case 1:
		operationUnit = calculator.Multiplier
	case 2:
		operationUnit = calculator.SquareMultiplier
	}
	// Operate라는 기능을 다양한 동작으로 수행할 수 있다.
	operationResult, operationErr = operationUnit.Operate(inputs...)
	// result
	if operationErr != nil {
		fmt.Println("[Error]", operationErr)
	} else {
		fmt.Println("Result:", operationResult)
	}
}

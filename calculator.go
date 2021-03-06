package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("______C A L C U L A T O R__________")
	c := Calc{}
	c.Operate(splitInfo(readInput()))
}

type Calc struct{}

func (Calc) Operate(operador1 int, operador2 int, operador string) int {
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "Not supported")
		return 0
	}
}

func parseInfo(valuesint []string, operator string) (int, int, string) {
	i1, _ := strconv.Atoi(valuesint[0])
	i2, _ := strconv.Atoi(valuesint[1])
	return i1, i2, operator
}

func readInput() string {
	fmt.Println("Please enter your operation:")
	var input string
	fmt.Scanln(&input)
	return input
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func splitInfo(entrada string) (int, int, string) {
	if strings.Contains(entrada, "+") {
		arrResult := strings.Split(entrada, "+")
		return parseInfo(arrResult, "+")
	}
	if strings.Contains(entrada, "*") {
		arrResult := strings.Split(entrada, "*")
		return parseInfo(arrResult, "*")
	}
	if strings.Contains(entrada, "-") {
		arrResult := strings.Split(entrada, "-")
		return parseInfo(arrResult, "-")
	}
	if strings.Contains(entrada, "/") {
		arrResult := strings.Split(entrada, "/")
		return parseInfo(arrResult, "/")
	}
	return 0, 0, "INVALID"
}

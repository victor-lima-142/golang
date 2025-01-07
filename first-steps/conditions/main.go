package conditions

import "fmt"

/* If- Else Statement funcs */

/* If Statement funcs */
func HigherThanTen(n int8) {
	if n > 10 {
		fmt.Printf("%d is higher than 10\n", n)
	}
	fmt.Printf("%d is lower than 10\n", n)
}

/* If- Else If Statement funcs */
func HigherThanHundred(n int8) {
	if n > 100 {
		fmt.Printf("%d is higher than 100\n", n)
	} else {
		fmt.Printf("%d is lower than 100\n", n)
	}
}

/* If- Else If- Else Statement funcs */
func HigherThanHundredAndTen(n int8) {
	if n > 100 {
		fmt.Printf("%d is higher than 100\n", n)
	} else if n > 10 {
		fmt.Printf("%d is higher than 10\n", n)
	} else {
		fmt.Printf("%d is lower than 10\n", n)
	}
}

/* Switch Statement funcs */
func SwitchStatement(symbol string) {
	switch symbol {
	case "+":
		fmt.Println("Addition")
	case "-":
		fmt.Println("Subtraction")
	case "*":
		fmt.Println("Multiplication")
	case "/":
		fmt.Println("Division")
	default:
		fmt.Println("Unknown operation")
	}
}

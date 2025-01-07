package types

import "fmt"

/* String type */
func StringType() {
	var name string = "Gopher"
	fmt.Println("I'm a string:", name)
}

/* Integer type */
func IntegerType() {
	var age int8 = 10
	var process int16 = 1000
	var total int32 = 100000
	var balance int64 = 1000000000
	var bigBalance int = 1000000000000000000
	var negative int = -1000
	var unsigned uint = 1000
	fmt.Println("I'm an integer (int8):", age)
	fmt.Println("I'm an integer (int16):", process)
	fmt.Println("I'm an integer (int32):", total)
	fmt.Println("I'm an integer (int64):", balance)
	fmt.Println("I'm an integer (int):", bigBalance)
	fmt.Println("I'm an integer (negative):", negative)
	fmt.Println("I'm an unsigned integer (positive only):", unsigned)
}

/* Float type */
func FloatType() {
	var price float32 = 10.50
	var total float64 = 100.50
	fmt.Println("I'm a float (float32):", price)
	fmt.Println("I'm a float (float64):", total)
}

/* Boolean type */
func BooleanType() {
	var isActive bool = true
	fmt.Println("I'm a boolean:", isActive)
}

/* Array type */
func ArrayType() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("I'm an array:", numbers)
}

/* Slice type */
func SliceType() {
	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println("I'm a slice:", numbers)
	numbers = append(numbers, 6)
	fmt.Println("I'm a slice (after append):", numbers)
}

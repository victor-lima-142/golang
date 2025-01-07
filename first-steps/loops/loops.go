package loops

import "fmt"

/* For Loop funcs */
func SimpleForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

/* For Loop with Range funcs */
func ForLoopWithRange(n int8) {
	var i int8
	for i = 0; i < n; i++ {
		fmt.Println(i)
	}
}

/* For Loop with Range and Continue funcs */
func ForLoopOnlyOdd(n int8) {
	var i int8
	for i = 0; i < n; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

/* For Loop with Range and Break funcs */
func ForLoopBreakAtFive(n int8) {
	var i int8
	for i = 0; i < n; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

/* For in Loop funcs */
func ForInLoop() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i, number := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, number)
	}
}

package main

import (
	"first-steps/conditions"
	"first-steps/hello"
	"first-steps/loops"
	"fmt"
)

func main() {
	// Call the HelloWorld function from the hello package
	fmt.Println("Hello World:")
	hello.HelloWorld()

	// Call the HigherThanTen function from the conditions package
	fmt.Println("HigherThanTen(5):")
	conditions.HigherThanTen(5)

	// Call the HigherThanHundred function from the conditions package
	fmt.Printf("HigherThanHundred(125):\n")
	conditions.HigherThanHundred(125)

	// Call the HigherThanHundredAndTen function from the conditions package
	fmt.Println("HigherThanHundredAndTen(50):")
	conditions.HigherThanHundredAndTen(50)

	// Call the SwitchStatement function from the conditions package
	fmt.Println("SwitchStatement(\"+\"):")
	conditions.SwitchStatement("+")

	// Call the SimpleForLoop function from the loops package
	fmt.Println("SimpleForLoop:")
	loops.SimpleForLoop()

	// Call the ForLoopWithRange function from the loops package
	fmt.Println("ForLoopWithRange(5):")
	loops.ForLoopWithRange(5)

	// Call the ForLoopOnlyOdd function from the loops package
	fmt.Println("ForLoopOnlyOdd(10):")
	loops.ForLoopOnlyOdd(10)

	// Call the ForLoopBreakAtFive function from the loops package
	fmt.Println("ForLoopBreakAtFive(10):")
	loops.ForLoopBreakAtFive(10)

	// Call the ForInLoop function from the loops package
	fmt.Println("ForInLoop:")
	loops.ForInLoop()
}

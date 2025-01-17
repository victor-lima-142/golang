package main

import (
	"fmt"
	"os"
	"strconv"
	"structures/basic"
)

func runQueue() {
	println("Running queue...")
	var queue *basic.Queue[int] = basic.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		println(item)
	}
}

func runStack() {
	println("Running stack...")
	var stack *basic.Stack[int] = basic.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		println(item)
	}
}

func runHashMap() {
	println("Running hash map...")
	basic.HashMap()
}

func main() {
	var continueLoop = true
	for continueLoop {
		println("Which command you wanna run?\n[1] Queue\n[2] Stack\n[3] HashMap\n[4] Exit")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			println("Error: " + err.Error())
			continueLoop = false
			os.Exit(1)
		}
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			println("Error: " + err.Error())
			continueLoop = false
			os.Exit(1)
		}
		switch inputInt {
		case 1:
			runQueue()
			continueLoop = askIfContinue()
		case 2:
			runStack()
			continueLoop = askIfContinue()
		case 3:
			runHashMap()
			continueLoop = askIfContinue()
		case 4:
			println("Bye! :D")
			continueLoop = false
			os.Exit(0)
		default:
			println("Invalid command")
			continueLoop = false
			os.Exit(1)
		}
	}
}

func askIfContinue() bool {
	println("Continue? (y/n)")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		println("Error: " + err.Error())
		os.Exit(1)
	}
	response := input == "y" || input == "Y"
	if !response {
		println("Bye! :D")
	}
	return response
}

package main

import "structures/basic"

func runQueue() {
	var queue *basic.Queue[int] = basic.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		println(item)
	}
}

func main() {
	runQueue()
}

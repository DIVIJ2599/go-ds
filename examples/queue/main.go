package main

import (
	"fmt"

	"github.com/divijsehgal25/go-ds/queue"
)

func main() {
	fmt.Println("Unlimited Queue Example")

	unlimitedQueue := queue.New[int]()

	unlimitedQueue.Enqueue(10)
	unlimitedQueue.Enqueue(20)
	unlimitedQueue.Enqueue(30)

	fmt.Println("Queue values:", unlimitedQueue.Values())
	fmt.Println("Size:", unlimitedQueue.Size())

	frontValue, ok := unlimitedQueue.Front()
	if ok {
		fmt.Println("Front value:", frontValue)
	}

	value, ok := unlimitedQueue.Dequeue()
	if ok {
		fmt.Println("Dequeued value:", value)
	}

	fmt.Println("Queue values after dequeue:", unlimitedQueue.Values())
	fmt.Println("Size after dequeue:", unlimitedQueue.Size())

	fmt.Println()
	fmt.Println("Fixed Size Queue Example")

	fixedQueue := queue.NewWithSize[string](2)

	fmt.Println(fixedQueue.Enqueue("A"))
	fmt.Println(fixedQueue.Enqueue("B"))
	fmt.Println(fixedQueue.Enqueue("C"))

	fmt.Println("Is full:", fixedQueue.IsFull())
	fmt.Println("Size:", fixedQueue.Size())
	fmt.Println("Capacity:", fixedQueue.Capacity())

	firstValue, ok := fixedQueue.Dequeue()
	if ok {
		fmt.Println("Dequeued value:", firstValue)
	}

	secondValue, ok := fixedQueue.Dequeue()
	if ok {
		fmt.Println("Dequeued value:", secondValue)
	}

	thirdValue, ok := fixedQueue.Dequeue()
	if ok {
		fmt.Println("Dequeued value:", thirdValue)
	} else {
		fmt.Println("Queue is empty")
	}
}

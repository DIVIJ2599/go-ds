package main

import (
	"fmt"
	"github.com/divij2599/go-ds/stack"
)

func main() {
	fmt.Println("Unlimited Stack Example")

	unlimitedStack := stack.New[int]()

	unlimitedStack.Push(10)
	unlimitedStack.Push(20)
	unlimitedStack.Push(30)

	fmt.Println("Size:", unlimitedStack.Size())

	topValue, ok := unlimitedStack.Peek()
	if ok {
		fmt.Println("Top value:", topValue)
	}

	value, ok := unlimitedStack.Pop()
	if ok {
		fmt.Println("Popped value:", value)
	}

	fmt.Println("Size after pop:", unlimitedStack.Size())

	fmt.Println()
	fmt.Println("Fixed Size Stack Example")

	fixedStack := stack.NewWithSize[string](2)

	fmt.Println(fixedStack.Push("A")) 
	fmt.Println(fixedStack.Push("B")) 
	fmt.Println(fixedStack.Push("C"))

	fmt.Println("Is full:", fixedStack.IsFull())
	fmt.Println("Size:", fixedStack.Size())
	fmt.Println("Capacity:", fixedStack.Capacity())

	firstValue, ok := fixedStack.Pop()
	if ok {
		fmt.Println("Popped value:", firstValue)
	}

	secondValue, ok := fixedStack.Pop()
	if ok {
		fmt.Println("Popped value:", secondValue)
	}

	thirdValue, ok := fixedStack.Pop()
	if ok {
		fmt.Println("Popped value:", thirdValue)
	} else {
		fmt.Println("Stack is empty")
	}
}

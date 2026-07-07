package main

import "fmt"

func main() {
	myStack := Stack{}
	if myStack.IsEmpty() {
		fmt.Println("Stack is empty: ", myStack.IsEmpty())
	}
	myStack.Push(10)
	fmt.Println("")
	fmt.Printf("Pushed: 10\n")
	myStack.Push(20)
	fmt.Printf("Pushed: 20\n")
	myStack.Push(30)
	fmt.Printf("Pushed: 30\n")
	fmt.Printf("Size: %d\n\n", myStack.Size())

	top, _ := myStack.Peek()
	fmt.Printf("Peek: %d\n", top)
	fmt.Printf("Size: %d\n\n", myStack.Size())

	pop, _ := myStack.Pop()
	fmt.Printf("Pop: %d\n", pop)
	pop, _ = myStack.Pop()
	fmt.Printf("Pop: %d\n", pop)
	fmt.Printf("Size: %d\n\n", myStack.Size())
	pop, _ = myStack.Pop()
	fmt.Printf("Pop: %d\n", pop)
	if myStack.IsEmpty() {
		fmt.Printf("Stack is empty: %t\n\n", myStack.IsEmpty())
	}
	pop, err := myStack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Stack struct {
	Array []int
}

func (a *Stack) Push(value int) {
	a.Array = append(a.Array, value)
}
func (a *Stack) Pop() (removeVal int, err error) {
	if len(a.Array) == 0 {
		return 0, fmt.Errorf("Pop from empty stack: error: stack is empty")
	}
	removeVal = a.Array[len(a.Array)-1]
	a.Array = a.Array[:len(a.Array)-1]
	return removeVal, nil
}

func (a *Stack) Peek() (topVal int, err error) {
	if len(a.Array) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	topVal = a.Array[len(a.Array)-1]
	return topVal, nil
}

func (a *Stack) IsEmpty() bool {
	if len(a.Array) != 0 {
		return false
	}
	return true
}

func (a *Stack) Size() (lenStack int) {
	return len(a.Array)
}

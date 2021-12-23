package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Stack successfully created with a size: ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())

	stack.Push("Go")
	stack.Push(1999)
	stack.Push("alexandre")
	stack.Push(12.465)

	fmt.Println("Stack size: ", stack.Size())

	for !stack.Empty() {
		value, _ := stack.Pop()
		fmt.Println("Taking value out: ", value)
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Empty?", stack.Empty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Empty stack!")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}

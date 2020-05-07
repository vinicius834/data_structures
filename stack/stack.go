package stack

import "errors"

type Stack struct {
	Items []interface{}
	Size  int
}

type StackFunctions interface {
	Push(item interface{})
	Pop() error
	Peek() (interface{}, error)
	IsEmpty() bool
}

func (stack *Stack) Push(item interface{}) {
	stack.Items = append(stack.Items, item)
	stack.Size++
}

func (stack *Stack) Pop() error {
	if stack.IsEmpty() {
		return errors.New("Stack is empty")
	}
	stack.Items = stack.Items[0 : stack.Size-1]
	stack.Size--
	return nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return stack.Items[stack.Size-1], nil
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.Items) == 0
}

package stack_test

import (
	stck "data_structures/stack"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	dataTets := []struct {
		value        interface{}
		expected     interface{}
		expectedSize int
	}{
		{value: 1, expected: 1, expectedSize: 1},
		{value: "Vinicius", expected: "Vinicius", expectedSize: 2},
		{value: 2.0, expected: 2.0, expectedSize: 3},
	}

	stack := stck.Stack{}
	for _, item := range dataTets {
		stack.Push(item.value)

		assert.Equal(t, item.expectedSize, stack.Size)
	}
}

func TestPop(t *testing.T) {
	stack1 := stck.Stack{Items: []interface{}{1, "vinicius", 2.0}, Size: 3}
	stack2 := stck.Stack{Items: []interface{}{1, "vinicius"}, Size: 2}
	stack3 := stck.Stack{Items: []interface{}{1}, Size: 1}
	stack4 := stck.Stack{Items: []interface{}{}, Size: 0}
	dataTets := []struct {
		stack         stck.Stack
		expectedSize  int
		expectedStack stck.Stack
	}{
		{stack: stack1, expectedSize: 2, expectedStack: stck.Stack{Items: []interface{}{1, "vinicius"}, Size: 2}},
		{stack: stack2, expectedSize: 1, expectedStack: stck.Stack{Items: []interface{}{1}, Size: 1}},
		{stack: stack3, expectedSize: 0, expectedStack: stck.Stack{Items: []interface{}{}, Size: 0}},
		{stack: stack4, expectedSize: 0, expectedStack: stck.Stack{Items: []interface{}{}, Size: 0}},
	}

	for _, item := range dataTets {
		item.stack.Pop()

		assert.Equal(t, item.expectedSize, item.stack.Size)
		assert.Equal(t, item.expectedSize, len(item.stack.Items))
		assert.Equal(t, item.stack, item.expectedStack)
	}
}

func TestPeek(t *testing.T) {
	stackThreeItems := stck.Stack{Items: []interface{}{1, "vinicius", 2.0}, Size: 3}
	stackTwoItems := stck.Stack{Items: []interface{}{1, "vinicius"}, Size: 2}
	stackOneItems := stck.Stack{Items: []interface{}{1}, Size: 1}
	emptyStack := stck.Stack{Items: []interface{}{}, Size: 0}
	dataTets := []struct {
		stack         stck.Stack
		expectedItem  interface{}
		expectedError error
		expectedSize  int
	}{
		{stack: stackThreeItems, expectedItem: stackThreeItems.Items[stackThreeItems.Size-1], expectedError: nil, expectedSize: len(stackThreeItems.Items)},
		{stack: stackTwoItems, expectedItem: stackTwoItems.Items[stackTwoItems.Size-1], expectedError: nil, expectedSize: len(stackTwoItems.Items)},
		{stack: stackOneItems, expectedItem: stackOneItems.Items[stackOneItems.Size-1], expectedError: nil, expectedSize: len(stackOneItems.Items)},
		{stack: emptyStack, expectedItem: nil, expectedError: errors.New("Stack is empty"), expectedSize: len(emptyStack.Items)},
	}
	for _, item := range dataTets {
		peekedItem, err := item.stack.Peek()
		assert.Equal(t, item.expectedItem, peekedItem)
		assert.Equal(t, item.expectedError, err)
		assert.Equal(t, item.expectedSize, len(item.stack.Items))
	}
}

func TestIsEmpty(t *testing.T) {
	stackOneItems := stck.Stack{Items: []interface{}{1}, Size: 1}
	emptyStack := stck.Stack{Items: []interface{}{}, Size: 0}
	dataTets := []struct {
		stack          stck.Stack
		expectedAnswer bool
	}{
		{stack: stackOneItems, expectedAnswer: false},
		{stack: emptyStack, expectedAnswer: true},
	}
	for _, item := range dataTets {
		assert.Equal(t, item.expectedAnswer, item.stack.IsEmpty())
	}
}

func BenchmarkPush(b *testing.B) {
	stack := stck.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(b.N)
	}

}

func BenchmarkPop(b *testing.B) {
	stack := stck.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	stack := stck.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	stack := stck.Stack{}
	for i := 0; i < b.N; i++ {
		stack.IsEmpty()
	}

}

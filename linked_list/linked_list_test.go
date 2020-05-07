package linked_list_test

import (
	"data_structures/linked_list"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	data := []struct {
		node     *linked_list.Node
		expected error
	}{
		{node: &linked_list.Node{1, nil}, expected: nil},
		{node: &linked_list.Node{2, nil}, expected: nil},
		{node: &linked_list.Node{3, nil}, expected: nil},
		{node: nil, expected: errors.New("Node is nil")},
	}

	linkedList := linked_list.LinkedList{}
	for _, item := range data {
		err := linkedList.Insert(item.node)
		assert.Equal(t, item.expected, err)
	}
}

func TestGetByIndex(t *testing.T) {
	data := []struct {
		node          *linked_list.Node
		index         int
		expectedError error
	}{
		{node: &linked_list.Node{1, nil}, index: 0, expectedError: nil},
		{node: &linked_list.Node{2, nil}, index: 1, expectedError: nil},
		{node: &linked_list.Node{3, nil}, index: 2, expectedError: nil},
		{node: nil, index: -1, expectedError: errors.New("Not Found")},
	}

	linkedList := linked_list.LinkedList{}
	for index, item := range data {
		if index < len(data)-1 {
			linkedList.Insert(item.node)
		}

		_, err := linkedList.GetByIndex(item.index)
		assert.Equal(t, err, item.expectedError)
	}
}

func TestRemoveByIndex(t *testing.T) {
	third := &linked_list.Node{3, nil}
	second := &linked_list.Node{2, third}
	first := &linked_list.Node{1, second}

	linkedList := linked_list.LinkedList{first, 3}
	linkedListOneElement := linked_list.LinkedList{third, 1}
	data := []struct {
		linkedList *linked_list.LinkedList
		index      int
		expected   error
	}{
		{linkedList: &linkedList, index: 2, expected: nil},
		{linkedList: &linkedListOneElement, index: 0, expected: nil},
		{linkedList: nil, index: -1, expected: errors.New("Out of bound")},
	}

	for _, item := range data {
		err := item.linkedList.RemoveByIndex(item.index)
		assert.Equal(t, item.expected, err)
	}
}

func TestIsEmpty(t *testing.T) {
	data := []struct {
		linked_list *linked_list.LinkedList
		expected    bool
	}{
		{linked_list: &linked_list.LinkedList{&linked_list.Node{1, nil}, 1}, expected: false},
		{linked_list: &linked_list.LinkedList{}, expected: true},
	}

	for _, item := range data {
		isEmpty := item.linked_list.IsEmpty()
		assert.Equal(t, item.expected, isEmpty)
	}
}

func TestIndexOf(t *testing.T) {
	data := []struct {
		node          *linked_list.Node
		expected      int
		expectedError error
	}{
		{node: &linked_list.Node{1, nil}, expected: 0, expectedError: nil},
		{node: &linked_list.Node{2, nil}, expected: 1, expectedError: nil},
		{node: &linked_list.Node{3, nil}, expected: 2, expectedError: nil},
		{node: &linked_list.Node{-1, nil}, expected: -1, expectedError: errors.New("Not Found")},
		{node: nil, expected: -1, expectedError: errors.New("Node is nil")},
	}

	linkedList := linked_list.LinkedList{}
	for index, item := range data {
		if index < len(data)-2 {
			linkedList.Insert(item.node)
		}

		index, err := linkedList.IndexOf(item.node)
		assert.Equal(t, item.expected, index)
		assert.Equal(t, item.expectedError, err)
	}
}

func BenchmarkInsert(b *testing.B) {
	linkedList := linked_list.LinkedList{}
	for i := 0; i < b.N; i++ {
		linkedList.Insert(linked_list.NewNode(i, nil))
	}
}

func BenchmarkRemoveByIndex(b *testing.B) {
	linkedList := linked_list.LinkedList{}
	times := b.N
	for i := 0; i < b.N; i++ {
		linkedList.Insert(linked_list.NewNode(i, nil))
	}

	b.StartTimer()
	for i := 0; i < times; i++ {
		linkedList.RemoveByIndex(i)
	}
	b.StopTimer()
}

func BenchmarkGetByIndex(b *testing.B) {
	linkedList := linked_list.LinkedList{}
	times := b.N
	for i := 0; i < b.N; i++ {
		linkedList.Insert(linked_list.NewNode(i, nil))
	}

	b.StartTimer()
	for i := 0; i < times; i++ {
		linkedList.GetByIndex(i)
	}
	b.StopTimer()
}

func BenchmarkIsEmpty(b *testing.B) {
	linkedList := linked_list.LinkedList{}
	times := b.N
	for i := 0; i < times; i++ {
		linkedList.Insert(linked_list.NewNode(i, nil))
	}

	b.StartTimer()
	for i := 0; i < times; i++ {
		linkedList.IsEmpty()
	}
	b.StopTimer()
}

func BenchmarkIndexOf(b *testing.B) {
	linkedList := linked_list.LinkedList{}
	times := b.N
	for i := 0; i < b.N; i++ {
		linkedList.Insert(linked_list.NewNode(i, nil))
	}

	b.StartTimer()
	for i := 0; i < times; i++ {
		linkedList.IndexOf(linked_list.NewNode(i, nil))
	}
	b.StopTimer()
}

package linked_list

//Very slow. Just to demonstration. Don't use it in production

import (
	"errors"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

type LinkedListFunctions interface {
	Insert(newNode *Node) error
	GetByIndex() *Node
	RemoveByIndex(int)
	IsEmpty() bool
	IndexOf(node *Node) int
}

func NewNode(value interface{}, next *Node) *Node {
	return &Node{value, next}
}

func (list *LinkedList) Insert(newNode *Node) error {
	if newNode == nil {
		return errors.New("Node is nil")
	}
	if list.Head == nil {
		list.Head = newNode
		list.Size++
		return nil
	}

	temp := list.Head
	for {
		if temp.Next == nil {
			temp.Next = newNode
			list.Size++
			return nil
		}
		temp = temp.Next
	}
}

func (list *LinkedList) GetByIndex(indexToSearch int) (*Node, error) {
	node := list.Head
	for index := 0; index <= indexToSearch; index++ {
		if index == indexToSearch {
			return node, nil
		}
		node = node.Next
	}

	return nil, errors.New("Not Found")
}

func (list *LinkedList) RemoveByIndex(indexToRemove int) error {
	if indexToRemove < 0 || list.Size < indexToRemove {
		return errors.New("Out of bound")
	}

	index := 0
	node := list.Head
	previous := node

	if indexToRemove == 0 || list.Size == 1 {
		list.Head = node.Next
		node = nil
		list.Size--
		return nil
	}

	for index < indexToRemove {
		if node.Next == nil {
			return errors.New("Not Found")
		}
		previous = node
		node = node.Next
		index++
	}
	previous.Next = node.Next
	node = nil
	list.Size--
	return nil
}

func (list *LinkedList) IsEmpty() bool {
	return list.Size == 0
}

func (list *LinkedList) IndexOf(node *Node) (int, error) {
	if node == nil {
		return -1, errors.New("Node is nil")
	}
	index := 0
	nodeToCompare := list.Head
	for {
		if nodeToCompare.Value == node.Value {
			return index, nil
		}
		if nodeToCompare.Next == nil {
			return -1, errors.New("Not Found")
		}
		nodeToCompare = nodeToCompare.Next
		index++
	}
}

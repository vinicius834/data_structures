package queue

import "errors"

type Queue struct {
	Items []interface{}
}

type QueueFunctions interface {
	EnQueue(item interface{}) error
	DeQueue() (interface{}, error)
	IsEmpty() bool
	Peek() (interface{}, error)
}

func (queue *Queue) EnQueue(item interface{}) error {
	if item == nil {
		return errors.New("Item is nil")
	}
	queue.Items = append(queue.Items, item)
	return nil
}

func (queue *Queue) DeQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := queue.Items[0]
	queue.Items = queue.Items[1:len(queue.Items)]
	return item, nil
}

func (queue *Queue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return queue.Items[0], nil
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.Items) == 0
}

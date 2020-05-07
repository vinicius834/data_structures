package queue_test

import (
	"data_structures/queue"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnQueue(t *testing.T) {
	data := []struct {
		value             interface{}
		queueSizeExpected int
		expectedError     error
	}{
		{value: 1, queueSizeExpected: 1, expectedError: nil},
		{value: "Vinicius", queueSizeExpected: 2, expectedError: nil},
		{value: nil, queueSizeExpected: 2, expectedError: errors.New("Item is nil")},
	}

	queue := queue.Queue{}
	for _, item := range data {
		err := queue.EnQueue(item.value)
		assert.Equal(t, item.expectedError, err)
		assert.Equal(t, item.queueSizeExpected, len(queue.Items))
	}
}

func TestDeQueue(t *testing.T) {
	item1 := 1
	item2 := 2
	item3 := 3
	queueWithElements := queue.Queue{[]interface{}{item1, item2, item3}}
	queueWithoutElements := queue.Queue{[]interface{}{}}
	data := []struct {
		queue             queue.Queue
		queueSizeExpected int
		expectedError     error
		expectedItem      interface{}
	}{
		{queue: queueWithElements, queueSizeExpected: 2, expectedError: nil, expectedItem: item1},
		{queue: queueWithoutElements, queueSizeExpected: 0, expectedError: errors.New("Queue is empty"), expectedItem: nil},
	}

	for _, item := range data {
		itemDeQueued, err := item.queue.DeQueue()
		assert.Equal(t, item.expectedError, err)
		assert.Equal(t, item.queueSizeExpected, len(item.queue.Items))
		assert.Equal(t, item.expectedItem, itemDeQueued)
	}
}

func TestPeek(t *testing.T) {
	item1 := 1
	item2 := 2
	item3 := 3
	queueWithElements := queue.Queue{[]interface{}{item1, item2, item3}}
	queueWithoutElements := queue.Queue{[]interface{}{}}
	data := []struct {
		queue         queue.Queue
		expectedItem  interface{}
		expectedError error
	}{
		{queue: queueWithElements, expectedItem: item1, expectedError: nil},
		{queue: queueWithoutElements, expectedItem: nil, expectedError: errors.New("Queue is empty")},
	}

	for _, item := range data {
		itemReturned, err := item.queue.Peek()
		assert.Equal(t, item.expectedError, err)
		assert.Equal(t, item.expectedItem, itemReturned)
		assert.Equal(t, item.expectedError, err)

	}
}

func TestIsEmpty(t *testing.T) {
	queueWithElements := queue.Queue{[]interface{}{1}}
	queueWithoutElements := queue.Queue{[]interface{}{}}
	data := []struct {
		queue          queue.Queue
		expectedAnswer bool
	}{
		{queue: queueWithElements, expectedAnswer: false},
		{queue: queueWithoutElements, expectedAnswer: true},
	}

	for _, item := range data {
		assert.Equal(t, item.expectedAnswer, item.queue.IsEmpty())
	}
}

func BenchmarkEnQueue(b *testing.B) {
	qe := queue.Queue{}
	times := b.N
	for i := 0; i < times; i++ {
		qe.Items = append(qe.Items, i)
	}
	b.StartTimer()
	for i := 0; i < times; i++ {
		qe.EnQueue(b.N)
	}
	b.StopTimer()
}

func BenchmarkDeQueue(b *testing.B) {
	qe := queue.Queue{}
	times := b.N
	for i := 0; i < times; i++ {
		qe.Items = append(qe.Items, i)
	}
	b.StartTimer()
	for i := 0; i < times; i++ {
		qe.DeQueue()
	}
	b.StopTimer()
}

func BenchmarkPeek(b *testing.B) {
	qe := queue.Queue{}
	times := b.N
	for i := 0; i < times; i++ {
		qe.Items = append(qe.Items, i)
	}
	b.StartTimer()
	for i := 0; i < times; i++ {
		qe.Peek()
	}
	b.StopTimer()
}

func BenchmarkIsEmpty(b *testing.B) {
	qe := queue.Queue{}
	for i := 0; i < b.N; i++ {
		qe.Items = append(qe.Items, i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		qe.IsEmpty()
	}
	b.StopTimer()
}

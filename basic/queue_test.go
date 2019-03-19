package basic

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}


func TestQueue_IsEmpty(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue.IsEmpty())
}


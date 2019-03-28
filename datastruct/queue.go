package datastruct


//数组队列
type Queue struct {
	data [100]int
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{
	}
}

func (q *Queue)Enqueue(data int) {
	if q.head == 100 {
		panic("queue is in full")
	}
	q.data[q.head] = data
	q.head++
}

func (q *Queue)Dequeue() int  {
	if q.tail == q.head {
		panic("queue is empty")
	}
	res :=  q.data[q.tail]
	q.tail++
	return res
}

func (q *Queue)Size() int {
	return q.head-q.tail+1
}

func (q *Queue) IsEmpty() bool  {
	return  q.head==q.tail
}
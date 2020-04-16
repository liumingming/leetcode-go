package datastruct


//数组栈
type Stack struct {
	capacity int
	data []int
	top int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity:capacity,
		data:make([]int, capacity, capacity),
	}
}

func (s *Stack)Push(data int) {
	s.data[s.top] = data
	s.top++
}

func (s *Stack)Pop() int  {
	if s.top == 0 {
		panic("stack is in bottom")
	}
	s.top--
	return s.data[s.top]
}

func (s *Stack)Size() int {
	return s.top
}

func (s *Stack) IsEmpty() bool  {
	return  s.top == 0
}
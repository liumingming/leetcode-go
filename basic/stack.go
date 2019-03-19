package basic


//数组栈
type Stack struct {
	data [100]int
	top int
}

func NewStack() *Stack {
	return &Stack{
	}
}

func (s *Stack)Push(data int) {
	if s.top == 100 {
		panic("stack is in full")
	}
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
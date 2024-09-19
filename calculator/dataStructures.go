package dataStructures


type Stack struct {
	array []int
}

func (s Stack) Push(value int) {
	append(s.array, value)
}

func (s Stack) Pop() int {
	if len(s.array) == 0 {
		panic("No elements in stack")
	}
	value := s.array[len(s.array) - 1]

}

func (s Stack) Length() int {
	return len(s.array)
}
package dataStructures


type Stack struct {
	array []int
}

func (s Stack) Push(value int) {
	append(s.array, value)
}

func (s Stack) Pop() (value int, err bool) {
	if len(s.array) == 0 {
		return (0, true)
	}
	value := s.array[len(s.array) - 1]
	return (value, false)
}

func (s Stack) Length() int {
	return len(s.array)
}
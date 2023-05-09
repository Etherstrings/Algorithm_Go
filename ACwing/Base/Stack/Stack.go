package Stack

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(element interface{}) {

	s.data = append(s.data, element)

}

func (s *Stack) Len() int {

	return len(s.data)

}

func (s *Stack) Peek() interface{} {

	return s.data[len(s.data)-1]

}

func (s *Stack) Pop() interface{} {

	if len(s.data) == 0 {

		return nil

	}

	res := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return res

}

func main() {

	s := Stack{}

	s.Push("golang")

	s.Push("stack")

	s.Push(100)

	fmt.Println(s.Pop()) //"100"

	fmt.Println(s.Pop()) //"stack"

	fmt.Println(s.Pop()) //"golang"

}

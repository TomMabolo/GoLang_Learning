package main

import "fmt"

type Stack struct {
	node int
	data [10]int 
}

func (stack *Stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}

func (stack *Stack) Pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

func main() {
	s := Stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	fmt.Println(s)

	for index := 1; index < 11; index++ {
		fmt.Println(s.Pop())
	}
}
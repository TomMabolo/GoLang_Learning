package main

import (
	"fmt"
	"strconv"
)


func SortNumber(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}

func FoundError() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
}

type stack struct {
	node int
	data [10]int
}

func (stack *stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}

func (stack *stack) Pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

func (stack *stack) ToString() string {
	var s string
	for i := 0; i < stack.node; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s += "[" + k + ":" + v + "] "

	}

	return s
}

func Tooargs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}

}

func Gblq(number int) []int {
	s := make([]int, number)
	for i := 1; i <= number; i++ {
		if i < 3 {
			s[i-1] = 1
		} else {
			s[i-1] = s[i-2] + s[i-3]

		}

	}
	return s

}

func ArrayMap(numbers []int, f func(int) int) []int {
	for i, v := range numbers {
		numbers[i] = f(v)
	}

	return numbers
}

func MaxOf(numbers []int) int {
	var max = 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max

}

func MinOf(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min

}

func ArrSort(numbers []int) {
	var temp int
	len := len(numbers)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if numbers[i] > numbers[j] {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
}

func PlusTwo() func(int) int {
	v := func(number int) int {
		return number + 2
	}
	return v
}

func PlusX(x int) func(int) int {
	v := func(y int) int {
		return x + y
	}
	return v

}

func main() {
	a, b := SortNumber(7, 2)
	fmt.Println(a, b)

	s := stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	fmt.Println(s)
	for index := 1; index < 11; index++ {
		fmt.Println(s.Pop());
	}

	t := stack{}
	for index := 1; index < 11; index++ {
		t.Push(index)
	}
	fmt.Println(t.ToString())

	Tooargs(1, 2, 3, 4)

	s1 := Gblq(20)
	fmt.Println(s1)

	arr := []int{1, 3, 4}
	mapArr := ArrayMap(arr, func(arg1 int) int {
		return arg1 * arg1
	})
	fmt.Println(mapArr)

	arr1 := []int{9, 3, 4, 7, 11, 2}
	max := MaxOf(arr1)
	fmt.Println(max)

	arr2 := []int{9, 3, 4, 7, 11, 2}
	min := MinOf(arr2)
	fmt.Println(min)

	arr3 := []int{9, 3, 4, 7, 11, 2}
	ArrSort(arr3)
	fmt.Println(arr3)

	fmt.Println(PlusTwo()(3))
	f := PlusTwo()
	fmt.Println(f(3))

	f1 := PlusX(7)
	fmt.Println(f1(3))
}

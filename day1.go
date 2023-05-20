package main

import (
	"fmt"
	"unicode/utf8"
)

func First() {
	fmt.Println("first func start")
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	fmt.Println("first func end")
}

func Second() {
	fmt.Println("second func start")
	i := 1
HERE:
	fmt.Println(i)
	if i < 10 {
		i++
		goto HERE
	}
	fmt.Println("second func end")
}

func Third() {
	fmt.Println("third func start")
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("The%dtitle%d\n", i+1, v)
	}
	fmt.Println("third func end")
}

func FizzBuzz(number int) {
	fmt.Println("FizzBuzz func start")
	for i := 1; i <= number; i++ {
		switch true {

		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s ", "FizzBuzz")
		case i%3 == 0:
			fmt.Printf("%s ", "Fizz")
		case i%5 == 0:
			fmt.Printf("%s ", "Buzz")
		default:
			fmt.Printf("%d ", i)
		}

	}
	fmt.Println()
	fmt.Println("FizzBuzz func end")

}

func PrintA(number int) {
	a := "A"
	for i := 0; i < number; i++ {
		fmt.Println(a)
		a = a + "A"
	}
}

func CountStr() {
	str := "asSASA ddd dsjkdsjs dk"
	strRune := []rune(str)
	fmt.Printf("len:%d\n", len(str))
	fmt.Printf("rune:%d\n", len(strRune))
	fmt.Printf("RuneCountInString:%d\n", utf8.RuneCountInString(str))
}

//ChangeStr 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。
func ChangeStr() {
	str := "asSASA ddd dsjkdsjs dk"
	strRune := []rune(str)
	// strRune[4] = 'a'
	// strRune[5] = 'b'
	// strRune[6] = 'c'
	copy(strRune[4:4+3], []rune("abc"))
	cStr := string(strRune)

	fmt.Println(cStr)
}

//ReverseStr
func ReverseStr(str string) {
	rs := []rune(str)
	var newStr string
	for i := len(rs) - 1; i >= 0; i-- {

		newStr = newStr + string(rs[i])
	}
	fmt.Println(newStr)
	var newStr2 []rune

	for i := len(rs) - 1; i >= 0; i-- {
		newStr2 = append(newStr2, rs[i])
	}
	fmt.Println(string(newStr2))

}

// AverageSlice 编写计算一个类型是 float64 的 slice 的平均值的代码。
func AverageSlice() {
	s := []float64{103.44, 20.07, 14.11, 1.01}
	var total float64
	for _, v := range s {
		total += v
	}
	fmt.Println(total / float64(len(s)))
}

func main(){
	First()
	Second()
	Third()
	FizzBuzz(100)
	PrintA(100)
	CountStr()
	ChangeStr()
	ReverseStr("foobar")
	AverageSlice()
}
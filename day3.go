package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type base interface {}

func typeSwitch(f base) base {
	switch f.(type) {
	case int:
		return  f.(int)*2;
	case string:
		return f.(string)+f.(string)+f.(string)
	}
	return  f
}

func Map(arr []base,f func( base) base)[]base{
	for k,v :=range arr{
		arr[k] = f(v)
	}

	return  arr
}

func Cat(fileName string,showLine bool){
	file,err := os.OpenFile(fileName,os.O_RDWR,0666)
	if err!=nil{
		fmt.Println("open file err",err)
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	lineNumber := 1
	for {
		line,err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if showLine{
			fmt.Println(strconv.Itoa(lineNumber)+" "+line)
		}else {
			fmt.Println(line)
		}
		lineNumber++
		if err != nil{
			if err == io.EOF {
				break
			}else {
				fmt.Println("open file err",err)
				return
			}

		}
	}

}


func main(){
	numbers := []base{1,2,3,4}
	strings := []base{"2","b","c"}
	newNumbers := Map(numbers, typeSwitch)
	newStrings := Map(strings, typeSwitch)
	fmt.Println(numbers)
	fmt.Println(strings)
	fmt.Println(newNumbers)
	fmt.Println(newStrings)
	Cat("day3.txt",true)
}



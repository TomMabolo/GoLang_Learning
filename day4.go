package main

import (
	"fmt"
)

func First(){
	/*
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=1 ;i<11;i++{
		go func(number int) {
			fmt.Println(number)
		}(i)
	wg.Done()
	}
	wg.Wait()
	*/
	c := make(chan int)

	flag := make ( chan bool )
	go func(ch chan int,f chan bool) {
		for  {
			select {
			case <-flag :
				break
			case number := <-ch:
				fmt.Println(number)
			}
		}
	}(c,flag)

	for i:=1;i<11;i++{
		c<-i
	}

	flag<-false

	//close(c)

}

func  dup3(in <-chan  int) (<-chan int,<-chan int,<-chan int)  {
	a,b,c := make(chan int ,2) ,make(chan int ,2),make(chan int ,2)
	go func() {
		for  {
			x := <-in
			a<-x
			b<-x
			c<-x
		}
	}()
	return  a,b,c
}

func fib() <-chan int{
	x := make(chan int,2)
	a ,b ,c := dup3(x)
	go func() {
		x<-0
		x<-1
		<-a
		for  {
			x<- <-a+<-b
		}
	}()
	return  c
}

func Runfib()  {
	ch := fib()
	for i:=0;i<20 ; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	First()
	Runfib()
}


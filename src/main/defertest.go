package main

import (
	"fmt"
)

func main(){
	fmt.Println("test1=",test1())
	fmt.Println("test2=",test2())
}

func change(s ...int){
	fmt.Println()
	s = append(s,10)

}

func test1() int{
	var a int
	a =5
	defer func(){
		a++
	}()

	return a
}

func test2() (r int){
	a := 5
	r = a
	defer func(){
		r++
	}()
	return r
}

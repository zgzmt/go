package main

import "fmt"

func main(){
	ch := make(chan string)
	go func() {
		for i := 0; i< 10; i++{
			fmt.Println("i=",i)
		}
		ch <- "hello"
		fmt.Println("len=",len(ch),"cap=",cap(ch))
	}()
	str := <-ch
	fmt.Println(str)

	fmt.Println("=================")
	arry := []int{1,2,3,4,5,6}
	s := arry[0:3]
	s1 := append(s,1,2,3,4)
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(arry)
	s1[0] = 5
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(arry)

	s[1] = 6
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(arry)


}

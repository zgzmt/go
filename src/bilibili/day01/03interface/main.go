package main

import "fmt"

type speak interface{
	Sperk(string)string
}
type student struct{}
func (s *student) Sperk(things string)(talk string){
	if things == "sb"{
		talk = things
	}else{
		talk = "你好"
	}
	return
}

func main(){
	//var peo speak = student{}  不可以这样定义，这是错误的，原因是实现接口的是指针类型
	var peo speak
	stu := &student{}
	peo = stu
	var a int = 1
	fmt.Println(a)
	things := "bitch"
	fmt.Println(peo.Sperk(things))

}

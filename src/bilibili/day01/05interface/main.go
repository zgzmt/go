package main

import "fmt"

///接口嵌套与空接口
/*
空接口是没有定义类型的接口  type 接口名 interface{},任何类型都可以实现空接口
任何类型都可以赋值给空接口，空接口作为参数，表示该参数可以接口任何类型的值
空接口作为map中，表示map中的值也是可以是任何类型

 */

type sayer interface {
	say()
}
type mover interface {
	move()
}
type animal interface {
	sayer
	mover
}
type cat struct{}
func (c cat)say(){
	fmt.Printf("喵喵\n")
}
func (c cat)move(){
	fmt.Printf("走一步看看\n")
}
func main(){
	var c  cat
	c.move()
	c.say()
	var s sayer = c
	s.say()
	var a animal = c
	a.say()
	a.move()


}

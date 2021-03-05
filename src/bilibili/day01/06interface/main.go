package main

import "fmt"

/*
接口断言  x.(T),返回 x转换为T类型的值和一个bool值
空接口可以存储任何类型的值，要如何获取空接口存储的值，就需要断言
 */

func main(){
	var name interface{}
	name = "张三啊"
	stu,flag := name.(string)
	if flag == false{
		fmt.Printf("接口断言失败!\n")
		return
	}
	fmt.Printf("断言后的类型：%T,值：%s\n",stu,stu)
	name = false
	switch  name.(type){
	case string: fmt.Printf("是string类型\n")
	case int: fmt.Printf("是int类型\n")
	default:
		fmt.Printf("断言失败\n")
	}
}

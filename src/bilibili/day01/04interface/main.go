package main
/*
接口实现，结构体嵌套实现
*/
import "fmt"

type WashingMachine interface{
	dry()
	was()
}
type dryer struct{}
func (d dryer)dry(){
	fmt.Printf("甩一甩\n")
}
type haier struct{
	dryer
}
func (h haier)was(){
	fmt.Printf("洗一洗\n")
}

func main(){
	var h haier
	h.dry()
	h.was()
	h.dryer.dry()
	//var w WashingMachine = dryer{}  dryer没有实现接口的所有方法，所以不能这样赋值给接口
	var w1 WashingMachine = haier{}
	w1.dry()
	w1.was()
	var d = dryer{}
	d.dry()

}
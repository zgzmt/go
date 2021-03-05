package main

import "fmt"

type Phone interface {
	call()
	ans()
}

type huawei struct{
	name string
}
func (hw huawei)call(){
	fmt.Println(hw.name)
}
func (hw huawei)ans(){
	fmt.Println("ans()")
}
type iphone struct {
	name string
}
func (ip iphone)call(){
	fmt.Println(ip.name)
}
func (ip iphone)ans(){
	fmt.Println("ans()")
}
func main(){
	hw := huawei{"huawei"}
	ip := iphone{"iphone"}
	var phone Phone
	phone = hw
	phone.call()
	phone = ip
	phone.call()

}


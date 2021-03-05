package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
)

//D:\work\新打印中心梳理\mysql表关系（附加订单系统需求）v2.xlsx
func main(){
	_ ,err := excelize.OpenFile("D:\\work\\新打印中心梳理\\mysql表关系（附加订单系统需求）v2.xlsx")
	if err != nil {
		fmt.Print(err)
	}
}

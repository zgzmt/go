package main

import (
	"fmt"
	_ "github.com/mattn/go-oci8"
	"unsafe"
)

func main(){
	var temp1 int64
	temp1 = 32
	fmt.Printf("值：%d\n",unsafe.Sizeof(temp1))
	nums := []int{1,2,3,4}
	start := 0
	end := len(nums)-1
	result := 0
	for ;start < end; {
		result += min(nums[start],nums[start+1]) + min(nums[end],nums[end -1])
		start+=2
		end -=2
	}
	fmt.Println(result)

	//fmt.Println("数据库连接测试。。。")
	//db, err := sql.Open("oci8","wms/wms@wmstest.gzmpc.com/wms")
	//if err != nil {
	//	fmt.Println("连接错误"+err.Error())
	//}else{
	//	fmt.Println("成功")
	//}
	//defer db.Close()
}
func min(a,b int) int{
	if a > b {
		return b
	}else{
		return a
	}
}

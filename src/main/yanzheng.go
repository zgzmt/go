package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type student struct {
	a int64
	b struct{}
}
type nullstruct struct{}
type teacher struct {
	b struct{}
	a int64
}
type classroom struct {
	a int64
	b int32
}
type book struct{
	a int32
	b int64
}



func main(){
	var temp1 int64
	fmt.Println(unsafe.Sizeof(temp1))
	var temp2 struct{}
	fmt.Println(unsafe.Sizeof(temp2))
	temp3 := student{}
	temp4 := teacher{}
	fmt.Printf("temp3:%d,temp4:%d\n",unsafe.Sizeof(temp3),unsafe.Sizeof(temp4))
	temp5 :=student{1,nullstruct{},}
	fmt.Println("temp5:",unsafe.Sizeof(temp5))
	fmt.Printf("temp5地址：%p,%p\n",&temp5.a,&(temp5.b))
	temp6 := teacher{}
	fmt.Printf("temp6地址：%p,%p\n",&temp6.a,&(temp6.b))
	temp7 := classroom{}
	temp8 := book{}
	var a int32
	fmt.Printf("int32:%d,temp7:%d,temp8:%d\n",unsafe.Sizeof(a),unsafe.Sizeof(temp7),unsafe.Sizeof(temp8))

	arr := []int{1,2,3,4,5} //这是一个切片
	for i,v :=range arr {
		arr[i] = 0
		if i ==0 {
			arr[2] = 7
		}
		arr = append(arr,v)
		fmt.Printf("%d,",v)
	}
	fmt.Printf("\n%d\n",arr)
	arr2 := [5]int{1,2,3,4,5} //这是一个数据
	for i,v := range arr2 {
		if i == 0 {
			arr2[2] = 7
		}
		fmt.Printf("%d,",v)
	}
	fmt.Printf("\n%d\n",arr2)
	str := "abcderg"
	for _,v := range str{
		fmt.Printf(" %T  ",v)
			fmt.Println(v)
	}
	fmt.Println("end....")
	fmt.Printf("\rssaa\rwoijsjdf\n")
	fmt.Printf("%3d,%d\n",2133,2313)
	fmt.Printf("\r[%-10s] %5s\n","sws","wud")
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg1 *sync.WaitGroup){
		defer wg1.Done()
		fmt.Println("sleeping.....")
		time.Sleep(time.Second*3)
		fmt.Println("sleep end......")
	}(&wg)
	wg.Wait()
	fmt.Println("end...")
	//wg2 := sync.WaitGroup{}
	////wg4 := wg2
	//wg2.Add(1)
	//go func(wg3 sync.WaitGroup){
	//	defer wg3.Done()
	//	fmt.Println("sleeping wg3")
	//	time.Sleep(time.Second*3)
	//	fmt.Println("wg3 end ....")
	//}(wg2)
	//wg2.Wait()
	fmt.Println("end.....")

	var bbb byte
	fmt.Println(unsafe.Sizeof(bbb))
	//var ccc []byte
	ccc := [...]byte{'1','1','1','1',}
	ddd := ccc[:]
	fmt.Println(ccc)
	fmt.Println(ddd)
	fmt.Println(string(ddd))

	intb := []byte{1,1,1,1}
	bytebuffer := bytes.NewBuffer(intb)
	var temp int32
	binary.Read(bytebuffer,binary.BigEndian,&temp)
	fmt.Println("转换后的整形：",temp)

	intbuffer := bytes.NewBuffer([]byte{})
	binary.Write(intbuffer,binary.BigEndian,&temp)
	fmt.Println(intbuffer.Bytes())
	fmt.Println(len(intbuffer.Bytes()))

	n := 2

	n |= n >> 2
	fmt.Println(n)
}


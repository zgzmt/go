package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	cli ,err := net.Dial("tcp","127.0.0.1:9090")
	if err != nil {
		fmt.Println("初始链接失败：",err)
		return
	}
	id := []byte{1,1,1,1}
	wn ,err :=cli.Write(id)
	if err != nil {
		fmt.Println("写数据错误：",err)
		return
	}
	fmt.Println("写了多少：",wn)
	recivebyte := make([]byte,1024)
	n,_ :=cli.Read(recivebyte)
	fmt.Println("接收到的数据：",string(recivebyte[:n]))
	time.Sleep(time.Second * 10)

}

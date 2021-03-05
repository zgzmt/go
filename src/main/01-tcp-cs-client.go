package main

import (
	"fmt"
	"net"
)

func main(){
	client,err := net.Dial("tcp","127.0.0.1:8050")
	if err != nil {
		fmt.Println("net.Dial error :",err)
	}
	defer client.Close()
	str := "hello!"
	fmt.Println("客户端写数据")
	client.Write([]byte(str))
	fmt.Println("客户端读数据")
	buf := make([]byte,2048)
	client.Read(buf)
	fmt.Println(string(buf))

}

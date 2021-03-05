package main

import (
	"fmt"
	"net"
)

func main(){
	listener,err := net.Listen("tcp","127.0.0.1:8050")
	if err != nil {
		fmt.Println("net.Listen error :",err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器等待连接。。。。。")
	cnn,err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error :",err)
	}
	defer cnn.Close()
	for {
		addr := cnn.RemoteAddr()
		fmt.Println(addr.String(), "：连接成功，等待数据发送.....")
		buf := make([]byte, 2048)
		n, err := cnn.Read(buf)
		if err != nil {
			fmt.Println("cnn.Read error :", err)
		}
		fmt.Println(string(buf[:n]))
		cnn.Write(buf)
	}

}
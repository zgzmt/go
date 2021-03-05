package main

import (
	"fmt"
	"io"
	"net"
)

type clien struct {

}

func main(){
	listenserver ,err := net.Listen("tcp","127.0.0.1:8900")
	if err != nil {
		fmt.Println("listen err :",err)
		return
	}
	defer listenserver.Close()
	for {
		con ,err :=listenserver.Accept()
		if err != nil {
			fmt.Println("accept err :",err)
		}
		go handleCon(con)

	}
}
func handleCon(con net.Conn){
	defer con.Close()
	fmt.Println("是否是重连..")
	for {
		var acc []byte
		n ,err :=con.Read(acc)
		if err != nil{
			if err == io.EOF {
				fmt.Println("发送结束")
			}else{
				fmt.Println("read err:",err)
			}
		}
		fmt.Printf("读取了%d字节数据\n",n)
	}
}

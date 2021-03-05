package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServer(address string){
	//绑定监听地址
	listener , err := net.Listen("tcp",address)
	if err !=nil{
		log.Fatal(fmt.Sprintf("Listen err:%V",err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind:%s,start listening...",address))
	for{
		//accept阻塞等待连接建立或者listen中断才会返回
		conn , err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err:%v",err))
		}
		go Handle(conn)
	}
}
func Handle(conn net.Conn){
	//使用bufio提供的缓冲区功能
	reader := bufio.NewReader(conn)
	for{
		msg, err := reader.ReadString('\n')
		if err != nil{
			//遇到错误是连接终端或者关闭，io.Eof表示
			if err == io.EOF {
			    log.Println(fmt.Sprintf("connection closs"))
			}else{
				log.Println(err)
			}
		}
		log.Println(msg)
		b := []byte(msg)
		//将消息发回给客户端
		conn.Write(b)
	}
}

func main(){
	ListenAndServer(":8000")
	//signal.Notify()
}
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main(){
	c := make(chan os.Signal,1)
	go func(){
		fmt.Println("进入信号发送区...")
		signal.Notify(c,os.Kill,os.Interrupt)
		fmt.Println("退出信号发送区。。")
	}()
	fmt.Println("等待信号发送")
	time.Sleep(time.Second*3)
	signal.Stop(c)
	s := <- c
	fmt.Println("signal:",s)
}
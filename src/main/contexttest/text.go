package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"time"
)
func main(){
	//fmt.Printf("开始阶段,类型：%T,值：%v\n",context.Background())
	parenctx,_ := context.WithCancel(context.Background())
	fmt.Printf("withcancel获取：类型：%T,值:%v\n",parenctx,parenctx)
	timeoutctx ,_ := context.WithTimeout(parenctx,5*time.Second)
	var valuectx  context.Context
    msgchan :=  make(chan int)
    //msgchan <- 1
    var wait  sync.WaitGroup
	wait.Add(1)
	go func(ctx context.Context){
		fmt.Println("cancle携程")

		fmt.Printf("timeoutctx01 start run......\n")
		select{
		case <-ctx.Done():
			valuectx = context.WithValue(timeoutctx,"值",1)
			fmt.Printf("结束ctx:%d\n",ctx.Err())
			msgchan <- 1
		}
		wait.Done()
	}(timeoutctx)
	wait.Add(1)
    go func(ctx context.Context){
    	fmt.Println("valuectx 携程")
    	select{
    	case msg := <-msgchan:
			fmt.Printf("msgchan:%d\n",msg)
			fmt.Printf("valuectx 值:%d\n",valuectx.Value("值"))
		}
    	wait.Done()
	}(valuectx)
    fmt.Print("waitting other go end..................")
    wait.Wait()
    fmt.Println("end....")

}

package main

import (
	"fmt"
	"time"
)

/***
进度条模拟操作
进度条结构体要素：图形、进度条位置，总进度，当前位置，百分比
实现方法：初始化（）、进度条打印（）、结束（）、
***/
type bar struct{
	graph string //显示图形
	state string //进度条
	cur int32    //当前进度
	prece int32  //百分比
	tatla int32  //总大小
}
func (b *bar)NewBar(start,tatla int32){
	b.cur = start
	b.tatla = tatla
	b.graph = "#"
	b.prece = b.GetPrece()
}
func (b *bar)GetPrece()int32{
	presce := float32(b.cur)/float32(b.tatla)*100
	return int32(presce)
}
func (b *bar)Play(cur int32){
	b.cur = cur
	last := b.prece
	b.prece = b.GetPrece()
	if b.prece != last  {
		b.state = b.state+b.graph
	}
	fmt.Printf("\r[%-100s]%3d%%  %3d/%3d",b.state,b.prece,b.cur,b.tatla)
}
func (b *bar)Finsh(){
	fmt.Println()
}
func main(){
	var b bar
	b.NewBar(0,100)
	for i:=0 ;i <100; i++{
		time.Sleep(time.Second*1)
		b.Play(int32(i))
	}
	b.Finsh()
	fmt.Println(b.prece )
}


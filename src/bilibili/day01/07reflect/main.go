package main

import (
	"fmt"
	"reflect"
)

type myInt int

func reflectV(v interface{}){
	t := reflect.TypeOf(v)
	fmt.Printf("name:%v kind:%v\n",t.Name(),t.Kind())
}

func chageReflect(i interface{}){
	v := reflect.ValueOf(i)
	if v.Elem().Kind() == reflect.Int64{  //这里也要指针
//		v.SetInt(200) //错误用法，修改的是值拷贝
		v.Elem().SetInt(200)//真确用法，通过elem获取地址来修改
	}
}

func main(){
	var yourInt myInt
	var p *float32
	var r rune
	var yourInt1 int
	reflectV(yourInt)
	reflectV(p)
	reflectV(yourInt1)
	reflectV(r)
	type book struct{
		name string
	}
	b := book{
		"好姆雷特",
	}
	b1 := &book{
		"丧钟为谁而鸣",
	}
	var b2 *book
	reflectV(b)
	reflectV(b1)
	reflectV(b2)
	b2 = &book{
		"复活",
	}
	var a int64 = 10
	chageReflect(&a)
	fmt.Printf("a=%d",a)
}

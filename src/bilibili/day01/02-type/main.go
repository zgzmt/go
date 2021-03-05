package main

import (
	"encoding/json"
	"fmt"
)

type Student struct{
	Name string
	Age int
	Id int
}
type Class struct {
	Cla string
	Stu []*Student
}

func main(){
	c := &Class{
		Cla:"101",
		Stu:make([]*Student,0,100),
	}
	for i:=1; i<11; i++{
		stu := &Student{
			Name:fmt.Sprintf("stu%02d",i),
			Age:18,
			Id:i,
		}
		c.Stu = append(c.Stu,stu)
	}
	//json序列化
	data ,err := json.Marshal(c)
	if err != nil {
		fmt.Errorf("json序列化失败!")
		return
	}
	fmt.Printf("data=%s\n",data)

	data01 :=`{"Cla":"102","Stu":[{"Name":"stu11","Age":19,"Id":0},{"Name":"stu12","Age":20,"Id":2}]}`
	cls := &Class{}
	err = json.Unmarshal([]byte(data01),cls)
	if err != nil{
		fmt.Printf("方序列号出错:%s",err)
	}else {
		fmt.Printf("反序列化结果：%#v\n", cls)
		fmt.Printf("%s%v",cls.Cla,cls.Stu[1].Name)
	}
}

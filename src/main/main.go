package main

import (
	"fmt"
	"zipfile"

)

func main(){
	//fmt.Println("hello world!")
	//
	//slince := []int{1,2,3,4,85}
	//
	//slince2 := slince[1:3]
	//slince = append(slince, 6)
	//fmt.Println(slince2)
	//slince[2] = 5
	//
	//fmt.Println(slince2)
	//fmt.Println(slince2[0])
	//
	//map1 := make(map[string]string)
	//
	//value,exit := map1["a"]
	//fmt.Println("value",value)
	//fmt.Println("exit",exit)
	//
	//var a doc
	//a = 100
	//fmt.Println(a)
	zipfilepath := "Z:/2018.zip"
	textfilepath := "Z:/filepath.txt"
	str := zipfile.ReadTextFile(textfilepath)
	for _,s := range str{
		fmt.Println(s)
	}
	err := zipfile.ReadZipFile(zipfilepath,"Z:/beifen/",str)
	if err != nil {
		fmt.Println(err)
	}


}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	fithpath := `\\173.1.10.11/电脑室/物流系统升级/打印中心/中间库视图/`
	f ,err := os.Open(fithpath)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:",err)

	}
	fmt.Println(f.Name())
	fmt.Println()
	getFileLIst(fithpath)
}
func getFileLIst(path string)error {
	wf,ferr := os.Create("D:/filename.txt")
	defer wf.Close()
	var name string
	err := filepath.Walk(path,func(path string,f os.FileInfo,err error) error {


		if ferr != nil {
			fmt.Println("create err :",ferr)
		}

		if f == nil {
			return err
		}
		if f.IsDir() {
			fmt.Println("isDir:",f.Name())
		}else{
			fmt.Println("not dir:",f.Name())
			temp := strings.ReplaceAll(f.Name(),".sql",",")
			name += temp
		}
		return nil
	})
	wf.WriteString(name)
	return err
}


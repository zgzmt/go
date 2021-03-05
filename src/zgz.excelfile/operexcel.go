package zgz_excelfile

import (
	"fmt"
	"github.com/Luxurioust/excelize"
)

func Openexcel(excelpath string)*excelize.File{
	f , err :=excelize.OpenFile(excelpath)
	if err != nil {
		fmt.Printf("打开exccel错误,err:%s",err)
	}
	return f
}

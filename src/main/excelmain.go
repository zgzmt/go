package main

import (
	"database/sql"
	"fmt"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"mysqldb"
	_ "mysqldb"
	"os"
	"strconv"
	"strings"
)

func textmain() {
	xlsx, err := excelize.OpenFile("D:\\study\\Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	cell,_ := xlsx.GetCellValue("客户提货明细", "A2")
	fmt.Println(cell)
	// Get sheet index.
	index := xlsx.GetSheetIndex("客户提货明细")
	// Get all the rows in a sheet.
	rows,_:= xlsx.GetRows("sheet" + strconv.Itoa(index))
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
var db *sql.DB

//"user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
/*func init(){
	var err error;
	dsn := "root:iuap1businessdb3gzmpc@tcp(172.29.3.5:3306)/gzmpcli_db?charset=utf8mb4&parseTime=True"
	db,err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("数据库连接错误，error:%s",err)
	}else{
		fmt.Printf("open成功")
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping不通,error:%s",err)
	}else{
		fmt.Printf("ping通了")
	}
}*/
func main(){
	xlsx,err := excelize.OpenFile("D:\\study\\Workbook.xlsx")
	defer db.Close()
	if err != nil {
		fmt.Printf("打开excel错误，error:%s\n",err)
	}
	for _,sheetname := range xlsx.GetSheetList() {
		fmt.Printf("sheet name :%s\n",sheetname)
		rows ,_:= xlsx.GetRows(sheetname)
		i := 1
		for _,row := range rows{
			for _, celname := range row {
				var orgtablename ,tablecolumnname string
				tempstrs := strings.Split(celname,".")
				if len(tempstrs) > 1 {
					orgtablename = tempstrs[0]
					tablecolumnname = tempstrs[1]
				}else{
					orgtablename = "aaa"
					tablecolumnname = tempstrs[0]
				}
				prestr := strings.Split(orgtablename,"_")
				newtablename := strings.Replace(orgtablename,prestr[0],"li",1)
				fmt.Println("newtablename="+newtablename+"tablecolumnname="+tablecolumnname)
				checkflag := mysqldb.Selecttablenum(db,newtablename,tablecolumnname)
				if checkflag {
					xlsx.SetCellValue(sheetname,"B"+strconv.Itoa(i),newtablename)
					xlsx.SetCellValue(sheetname,"C"+strconv.Itoa(i),tablecolumnname)
					//xlsx.SetCe
				}
				i++
			}
		}
	}
	err = xlsx.Save()
	if err != nil {
		fmt.Println("写入excecl失败，err:",err)
	}
}

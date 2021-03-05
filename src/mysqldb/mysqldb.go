package mysqldb

import (
	"database/sql"
	"fmt"
	_"fmt"

)

type tablecolunm struct{
	table_name string
	colunm_name string
}


func Selecttablenum(db *sql.DB,tablename,columnname string)bool{
	var temptable tablecolunm
	sql := "select TABLE_NAME,column_name from information_schema.columns where column_name = ? and table_schema = 'gzmpcli_db' and TABLE_NAME = ?"
	err := db.QueryRow(sql,columnname,tablename).Scan(&temptable.table_name,&temptable.colunm_name)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	fmt.Println("查询结果："+temptable.colunm_name+" " +temptable.table_name)
	return true
}
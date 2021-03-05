package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	testhandlefunc := testhandlefunc
	http.HandleFunc("/test",testhandlefunc)
	http.ListenAndServe(":8060",nil)
}
func testhandlefunc(w http.ResponseWriter,r *http.Request){
	//var buf []byte
	defer r.Body.Close()
	fmt.Println("Method: ", r.Method)
	fmt.Println("URL: ", r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	fmt.Println("RemoteAddr: ", r.RemoteAddr)
	buf := make([]byte,1024)
	r.Body.Read(buf)
	result := string(buf)
	fmt.Println("结果：",result)
	//buf2 := []byte{'s','b'}
	buf2,err := json.Marshal(result)
	if err !=nil {
		fmt.Println("json err:",err)
	}
	w.Write(buf2)
}

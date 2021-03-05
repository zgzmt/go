package zipfile

import (
	"archive/zip"
	"fmt"
	"io"
	_ "io"
	"io/ioutil"
	"os"
	_ "os"
	"path/filepath"
	"strings"
	_ "strings"
)

func ReadZipFile(zipfilepath,destpath string,sourcefilename []string) error{
	reader,err := zip.OpenReader(zipfilepath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _,str := range sourcefilename {
		for _,file := range reader.File{
			//fmt.Println(file.Name)
			if str == file.Name {
				//fmt.Println(file.Name)
				destfilename := filepath.Join( destpath , file.Name)
				if file.FileInfo().IsDir() {
					err = os.MkdirAll(destfilename,os.ModePerm)
					if err != nil {
						return err
					}
				}else {
					if err = os.MkdirAll(filepath.Dir(destfilename),os.ModePerm); err != nil{
						return err
					}
					rc, err := file.Open()
					if err != nil {
						return err
					}
					defer rc.Close()
					outfile ,err := os.OpenFile(destfilename,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,os.ModePerm)
					if err != nil {
						return err
					}
					defer outfile.Close()
					_,err = io.Copy(outfile,rc)
					if err != nil {
						return  err
					}
				}
			}
		}
	}
	return nil
}

func ReadTextFile(fpath string)(result []string){
	date,err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	result = strings.Split(string(date),",")
	return result
}

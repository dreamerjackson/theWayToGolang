/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)


//自定义writer
type logWriter struct{}

func main(){
	resp,err := http.Get("http://tmall.com")

	if err !=nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	lw:= logWriter{}

	//读取resp.Body中的信息，写到os.Stdout中。os.Stdout实现了write接口，resp.Body实现了read接口。
	io.Copy(lw,resp.Body)
}

//自定义writer
func (logWriter) Write(bs []byte)(int,error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes",len(bs))
	return len(bs),nil
}

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

/*
go源代码中的接口：
http库中，get方法：

func Get(url string) (resp *Response, err error)
其中返回值Response：
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header     Header

	Body io.ReadCloser
}

io.ReadCloser:
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

*/

//不管是读取文件、网络等操作，都实现了相同的Reader接口。当传递[]byte进去，会将读取到的byte放置进去。成功会返回成功的个数。
package main

import (
"net/http"
"fmt"
"os"
)

func main(){
	resp,err := http.Get("https://tmall.com")

	if err !=nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	//设置大一点，read方法不会自动的扩容。
	bs:=make([]byte,99999)
	//read函数将读取到的数据放入到bs中。
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
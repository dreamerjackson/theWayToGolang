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

func main(){
	resp,err := http.Get("http://tmall.com")

	if err !=nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	//读取resp.Body中的信息，写到os.Stdout中。os.Stdout实现了write接口，resp.Body实现了read接口。
	io.Copy(os.Stdout,resp.Body)
}
/*
type Writer interface {
        Write(p []byte) (n int, err error)
}
*/

//io.copy源代码，读read接口，写入write接口

/*
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}

---------
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	if buf != nil && len(buf) == 0 {
		panic("empty buffer in io.CopyBuffer")
	}
	return copyBuffer(dst, src, buf)
}
---------
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	size := 32 * 1024
	if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
		if l.N < 1 {
			size = 1
		} else {
			size = int(l.N)
		}
	}
	if buf == nil {
		buf = make([]byte, size)//分配
	}
	for {   //疯狂读取并写入
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

*/
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"bufio"
	"unsafe"
)

var file *bufio.Writer
func memtoFile(pp *[]*[]*byte, deep int) {

	//最后一级指针时，开辟内存，存储密码。
	if deep == 1 {
		for i:=0;i<=9;i++{
			if (*pp)[i]!=nil{

				p:= (*string)(unsafe.Pointer((*pp)[i]))
				file.WriteString(*p+"\n")
			}
		}
		return
	}

	//刚开始传递11级指针的地址，deep=11递归下去，可以省略，改为传递11级指针，deep=10。
	if deep == 11 {
		memtoFile((*[]*[]*byte)(unsafe.Pointer(pp)), deep - 1);
		return
	}

	for i:=0;i<=9;i++{
		if (*pp)[i]!=nil{
			memtoFile((*[]*[]*byte)(unsafe.Pointer((*pp)[i])), deep - 1);
		}
	}
}
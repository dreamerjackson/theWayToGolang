/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "unsafe"

/*
最重要的递归函数，pp指针。deep为深度，str为字符串，password为密码
*/
func assign(pp *[]*[]*byte, deep int, str string,password string) {

	//最后一级指针时，开辟内存，存储密码。
	if deep == 1 {

		buf:= make([]*byte,10)
		(*pp)[getnum(str[10-deep])] = &buf
		p:= (*string)(unsafe.Pointer((*pp)[getnum(str[10-deep])]))
		*p = password
		return
	}

	//刚开始传递11级指针的地址，deep=11递归下去，可以省略，改为传递11级指针，deep=10。
	if deep == 11 {
		assign(pp, deep - 1, str,password);
		return
	}
	//判断该指针是否为空。如1131052403，当deep=10时，取出第一个数字1.判断pp[1]是否为Nil，为Nil就说明从来没有出现过第10位为1的qq号。
	//这时就会为pp[1]开辟10个指针大小的内存，初始化为空。
	//如果已经存在就继续递归。
	if (*pp)[getnum(str[10-deep])]!=nil{
		assign((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str,password);
	}else {
		buf:= make([]*byte,10)
		(*pp)[getnum(str[10-deep])] = &buf
		assign((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str,password);
	}
}
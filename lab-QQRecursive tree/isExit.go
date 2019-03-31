/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "unsafe"

/*
根据初始化递归树的原理，我们可以明白判断是否存在的意义。
如果在某一位出现了一个数字，其所在的指针为NULL，就说明在该位从来没有出现过这个数字。
相反，如果存在该数字，就首先说明其每一位的指针都不为空。

*/
var flag = true
var findresult string = ""
func isExit(pp *[]*[]*byte, deep int, str string) {

	if deep == 1 {
		if (*pp)[getnum(str[10-deep])]!=nil{
			findresult = *(*string)(unsafe.Pointer((*pp)[getnum(str[10-deep])]))
		}
		return
	}

	if deep == 11 {
		isExit(pp, deep - 1, str);
		return
	}

	if flag && (*pp)[getnum(str[10-deep])]!=nil{
		isExit((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str);
	}else {

		flag = false
		return
	}
}

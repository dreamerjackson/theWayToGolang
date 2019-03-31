/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "strings"

//字符转换为数字
func getnum(u uint8) uint8{
	return u - '0'
}

//qq补齐位数，判断是否为数字，字符转换为数字，数字不足补充0的算法
//获取QQ号 1131052403----qwerty
//截取数字10位，不足的补0，
//对于这个函数的改进，让我可以在查找qq函数时也可以用
func getQQ(s string) string{
	raw:= strings.Split(s,"----")[0]
	length := len(raw)
	if length < 10 {
		raw =  strings.Repeat("0",10-length) + raw
	}
	return raw
}

//判断qq全部为数字
func isAllNum(qq string ) bool{
	for _,ch := range qq{
		if ch < '0' || ch > '9'{
			return false
		}
	}
	return true
}


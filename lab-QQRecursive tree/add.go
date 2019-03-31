/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"log"
	"strings"
)

func add(qq string){


	qqslice := strings.Split(qq,"----")

	if len(qqslice) <2{
		log.Fatal("请输入正确的qq号，例如：1131052403----password")
		return
	}
	qqmian:= qqslice[0]
	qqpassword:= qqslice[1]


	qqlength := len(qqmian)

	if qqlength < 10 {
		qqmian =  strings.Repeat("0",10-qqlength) + qqmian
	}

	if isAllNum(qqmian)==false{
		log.Fatal("不是数字")
		return
	}

	pwdlength:= len(qqpassword)
	if pwdlength <=0{
		log.Fatal("没有设置密码")
		return
	}
	flag = true

	isExit(&g_pp,11,qqmian)

	if flag==true{
		log.Fatal("已经存在账号")
		return
	}

	assign(&g_pp,11,qqmian,qqpassword)

}

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"bufio"
	"os"
	"fmt"
)

//查找qq号
func search(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		flag = true
		findresult=""

		qq:= input.Text()
		qq= getQQ(qq)
		fmt.Println("搜索qq号：",qq)
		isExit(&g_pp,11,qq)
		if isAllNum(qq) && findresult!=""{
			fmt.Printf("结果为：%s\n",findresult)
		}
	}
}

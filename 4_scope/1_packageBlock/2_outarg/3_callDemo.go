/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main // main/4_scope.go

import (
	"fmt"
	"testdemo"
)
//调用另一个包中的函数和属性,注意大写字母为权限控制
func test() {

	testdemo.Haha()
	fmt.Println(testdemo.Birth)
}

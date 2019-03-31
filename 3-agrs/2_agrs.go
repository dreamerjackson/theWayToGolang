

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//变量的赋值
//变量是内存当中一段空间的抽象描述。变量的类型明确了空间的大小以及空间如何解析。
package main

import "fmt"



//1、外部变量声明
var i int
//2、外部连续声明
var U, V, W float64
//3、外部声明 赋值不带类型
var k = 0
//4、外部连续声明+赋值
var x, y float32 = -1, -2

//5、外部var括号内部
var (
	g       int
	u, v, s = 2.0, 3.0, "bar"

)


func main() {
	//函数内部的变量声明    声明的变量必须使用 否则报错
	var x string
	x = "Hello World"
	fmt.Println(x)

	//变量声明方式2  只限函数内部 自动推断类型
	 y := "jonson"
	fmt.Println(y)

	 //函数内部多样赋值
	var a,b string= "jonson","jackson"
	c,d := true,false
	e,f,g := "jonson",true,123
	fmt.Println(a,b,c,d,e,f,g)


	//特殊的类型声明
	type newstring = string
	var h newstring = "hello"
	fmt.Println(h)
}
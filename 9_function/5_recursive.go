/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//死循环

func  printforever(){
	fmt.Println("hello jonson")

	printforever()
}



//递归终止条件
var a =0
func  printf2(){
	fmt.Println("hello jonson")
	a++

	if a<3{
		printf2()
	}
}


//10进制转2进制
func  HexTobinary(a int){

	if a==0{
		return
	}

	b:= a%2
	a = a/2

	HexTobinary(a)

	fmt.Println(b)
}
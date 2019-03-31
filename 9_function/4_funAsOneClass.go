/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

//函数作为返回值
func makeGreeter() func() string{

	return func() string {
		return "hello jonson"
	}
}

//函数作为参数
func visit(numbers []int,callback func(int)){

	for _,n :=range numbers{
		callback(n)
	}
}


//函数作为变量
func main(){

	a:=makeGreeter()
	a()

}
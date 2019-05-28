/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
//指针作为函数参数,修改原来的值：
func main() {
	a := 10
	fmt.Printf("1、变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//10

	b := &a
	change(b)
	fmt.Printf("3、change函数调用之后，变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//20

	change0(a)
	fmt.Printf("5、change0函数调用之后，变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//20

}

func change(a *int) {
	fmt.Printf("2、change函数内，变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//20
	*a = 50
}

func change0(a int) {
	fmt.Printf("4、change0函数内，变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//20
	a = 90
}


//具有返回值的惯用写法，实现两个数据的交换
func swap0(x, y int) (int, int) {
	return y, x
}

//使用指针作为参数的写法
func swap(x, y *int) {
	*x, *y = *y, *x
}
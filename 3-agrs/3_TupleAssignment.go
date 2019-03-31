/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值。在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。这对于处理有些同时出现在元组赋值语句左右两边的变量很有帮助，例如我们可以这样交换两个变量的值：
//x, y = y, x
//
//a[i], a[j] = a[j], a[i]

//计算两个整数值的的最大公约数（GCD）
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main(){
	fmt.Println(gcd(21,6))
}
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"math"
)

//匿名函数
func main() {
	//无参匿名函数
	func() {
		fmt.Println("jonson")
	}()

	//有参匿名函数
	func(data int) {
		fmt.Println("data:", data)
	}(5)

	//有返回值的匿名函数
	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(9)

	fmt.Println("result:", result)

	//函数表达式
	greet := func() {
		fmt.Println("greet jonson")
	}
	greet()

	fmt.Printf("greet的类型是%T\n", greet)
}
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func greetjonson(){
	x:=0

	increment:= func() int{
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

//闭包与函数返回值
func makeEvenGenerator() func() int{
	i:=0
	return func() int{
		i+=2
		return i
	}
}
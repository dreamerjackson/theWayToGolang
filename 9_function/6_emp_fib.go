/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//0 1 1 2 3 5 8 13 21


//斐波拉契数列


func fib(index int) int{

	if index <2{
		return index
	}

	a:= 0
	b:= 1

	for i:=2;i<=index;i++{
		a,b = b,a+b
	}

	return b

}


func fib2(index int) int{

	if index <2{
		return index
	}

	return fib2(index-1) + fib2(index-2)
}


func fib3() func() int{

	a,b:=0,1

	return func() int {
		a,b = b,a+b
		return a
	}

}




func main(){

	fmt.Println(fib(4))
	fmt.Println(fib2(4))

	f:= fib3()

	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}

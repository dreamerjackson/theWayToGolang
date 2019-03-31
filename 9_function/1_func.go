/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

// 1、无参函数

func printstring(){
	fmt.Println("hello jonson")
}

// 2、带参函数

	func add(a,b int){
		fmt.Println("a+b=",a+b)
	}

// 3、返回值


func addres(a,b int) int{
	return a+b
}


// 4、有名字的返回值

func addresname(a,b int)(sum int){
	sum = a+b
	return
}


// 5、多返回值

func addmuti(a,b int)(int,int){
	add := a+b
	mul:= a*b

	return add,mul
}



// 6、不定个数参数


func addsum(nums ... int) int{
	var sum int
	for _,value :=range nums{
		sum +=value
	}
	return sum
}


/*

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
func name(parameter-list) (result-list) {
    body
}
形式参数列表描述了函数的参数名以及参数类型。这些参数作为局部变量，其值由参数调用者提供。返回值列表描述了函数返回值的变量名以及类型。如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任何值。在hypot函数中,

func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3,4)) // "5"
x和y是形参名,3和4是调用时的传入的实数，函数返回了一个float64类型的值。返回值也可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为0。如果一个函数在声明时，包含返回值列表，该函数必须以return语句结尾，除非函数明显无法运行到结尾处。例如函数在结尾时调用了panic异常或函数中存在无限循环。
正如hypot一样，如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型。下面2个声明是等价的：
func f(i, j, k int, s, t string)
func f(i int, j int, k int,  s string, t string)

下面，我们给出4种方法声明拥有2个int型参数和1个int型返回值的函数.blank identifier(译者注：卽下文的_符号)可以强调某个参数未被使用。

func add(x int, y int) int   {return x + y}
func sub(x, y int) (z int)   { z = x - y; return}
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }


*/
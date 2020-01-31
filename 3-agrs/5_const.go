

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//常量一旦声明不能改变，并且常量必须赋予初始值。此代码无效`func main() {const x int}`

package main

import (
	"time"
	"fmt"
)

//推断类型
const (
	m  =  1
	n  =  2
)
//明确类型
const noDelay time.Duration = 0
const timeout = 5 * time.Minute

//如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式写法，对应的常量类型也一样的。例如：
const (
	a = 1
	b
	c = 2
	d
)

func main(){

  const k = 8
  fmt.Println(a,b,c,d)  // "1 1 2 2"
}

//常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。
//
//一个常量的声明语句定义了常量的名字，和变量的声明语法类似，常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。例如，常量比变量更适合用于表达像π之类的数学常数，因为它们的值不会发生变化：

/*
所有常量的运算都可以在编译期完成，这样可以减少运行时的工作，也方便其他编译优化。当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越界、任何导致无效浮点数的操作等。
常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex和unsafe.Sizeof。
因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度：
const IPv4Len = 4
// parseIPv4 parses an IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
	var p [IPv4Len]byte
	// ...
}

*/


/*

一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。在下面的代码中，time.Duration是一个命名类型，底层类型是int64，time.Minute是对应类型的常量。下面声明的两个常量都是time.Duration类型，可以通过%T参数打印类型信息：
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
*/


// think?
/*

package main

import (
    "fmt"
)

func main() {
    const (
        a = 0.1
        b = 0.2
        c = 0.3
    )
    fmt.Println(a+b == c)
}

___________________

package main

import (
    "fmt"
)

func main() {
    const (
        A = 0.1
        B = 0.2
    )
    a, b := A, B
    fmt.Println(A+B==a+b)
}

*/
// 参考：https://blog.golang.org/constants
// https://www.ardanlabs.com/blog/2014/04/introduction-to-numeric-constants-in-go.html
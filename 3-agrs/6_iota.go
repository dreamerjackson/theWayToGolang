
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
// 在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

/*
const (
	a = iota
	b = iota
	c = iota
	d = iota
)
等价于：
const (
	a = iota
	b
	c
	d
)
*/
//下面是一个更复杂的例子，每个常量都是1024的幂：
//const (
//	_ = 1 << (10 * iota)
//	KiB // 1024
//	MiB // 1048576
//	GiB // 1073741824
//	TiB // 1099511627776             (exceeds 1 << 32)
//	PiB // 1125899906842624
//	EiB // 1152921504606846976
//	ZiB // 1180591620717411303424    (exceeds 1 << 64)
//	YiB // 1208925819614629174706176
//)
func main() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
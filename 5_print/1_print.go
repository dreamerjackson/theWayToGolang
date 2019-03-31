



/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */



//golang输出的格式化打印
package main

import "fmt"

func main(){


	//通用的格式
	str:="jonson"

	fmt.Printf("%T,%v\n",str,str)

	//布尔
	var booeanl  = true
	fmt.Printf("%T,%t\n",booeanl,booeanl)

	//特殊字符

	fmt.Printf("%%\n")
	fmt.Printf("\"\n")
	fmt.Printf("\\\n")


	//整数

	fmt.Printf("%T,%d\n",123,123)
	fmt.Printf("%T,%6d\n",123,123) //6代表长度
	fmt.Printf("%T,%06d\n",123,123)//0代表填充0
	fmt.Printf("%T,%b\n",123,123) //二进制
	fmt.Printf("%T,%o\n",123,123) //8进制
	fmt.Printf("%T,%x\n",123,123) //16进制
	fmt.Printf("%T,%#x\n",123,123) //前面加上0x
	fmt.Printf("%T,%#o\n",123,123)//前面加上0

	fmt.Printf("%T,%#X\n",123,123) //大写的X代表字母会大写

	fmt.Printf("% d,% d\n",123,-123)  //空格代表正数前方会预留一个空格

	// 浮点数
	fmt.Printf("%T,%f\n",123.456,123.456)
	fmt.Printf("%T,%10f\n",123.456,123.456)//长度
	fmt.Printf("%T,%.1f\n",123.456,123.456) //保留的小数位数
	fmt.Printf("%T,%10.2f\n",123.456,123.456)

	//字符串
	fmt.Printf("%T,%s\n","jonson","jonson")

	//字符串
	fmt.Printf("%T,%c\n",'c',97)
}
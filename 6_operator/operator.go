/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"



//算术运算符
func math(){


	a := 4
	b:=2

	fmt.Printf("a+b的结果为：%d\n", a+b)
	fmt.Printf("a-b的结果为：%d\n", a-b)
	fmt.Printf("a*b的结果为：%d\n", a*b)
	fmt.Printf("a/b的结果为：%d\n", a/b)

}

//关系运算符
func relation(){
	a := 4
	b := 2
	if(a==b){
		fmt.Printf("a与b相同\n")
	}else{
		fmt.Printf("a与b不同\n")
	}

	if(a<b){
		fmt.Printf("a小于b\n")
	}else{
		fmt.Printf("a大于b\n")
	}

	if(a>b){
		fmt.Printf("a大于b\n")

	}else{
		fmt.Printf("a小于b\n")
	}

	if(a<=b){
		fmt.Printf("a小于等于b\n")
	}else{
		fmt.Printf("a大于等于b\n")
	}

	if(a>=b){
		fmt.Printf("a大于等于b\n")

	}else{
		fmt.Printf("a小于等于b\n")
	}

}

//逻辑运算符
func logic(){
	a:=true
	b:=false

	if(a && b){
		fmt.Printf("a与b 同时为true\n")
	}else{
		fmt.Printf("a与b 不同时为true\n")
	}

	if(a || b){
		fmt.Printf("a与b 至少一个为true\n")
	}else{
		fmt.Printf("a与b 全部为false\n")
	}

	if(!b){
		fmt.Printf("取反成功\n")
	}else{
		fmt.Printf("取反失败\n")
	}

}

//位运算
func wei(){
	a := 3
	b:= 4

	fmt.Println("a & b :",a & b)
	fmt.Println("a | b :",a | b)
	fmt.Println("a ^ b :",a ^ b)
	fmt.Println("^ b :",^b)
	fmt.Println("a左移1位 :",a <<1)
	fmt.Println("a右移一位:",a >>1)


}

// 赋值运算符
func Assign(){
	a := 3
	var c int= 2

	c += a  // c =  c+a
	fmt.Println("c+=a的结果为:",c)

	c -= a //c = c -a
	fmt.Println("c-=a的结果为:",c)

	c *= a // c = a *c
	fmt.Println("c*=a的结果为:",c)

	c /= a  //   c =  c /a
	fmt.Println("c/=a的结果为:",c)
}
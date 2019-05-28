
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//切片基本声明与定义
package main

import "fmt"

//1、切片可以修改大小

//2、切片的拷贝不是单纯值的拷贝，一个切片指向了一个数组


//切片的声明1
var slice1 []int

//切片的声明2
var slice2 []int = make([]int,5)
var slice3 []int = make([]int,5,7)

func main(){
	//切片的声明3
	month:= [...]int{1,2,3,4,5,6,7,8,9}
	//切片的声明4
	month = [9]int{1:1,2:10}


	fmt.Println(month)
	//切片的声明5
	slice4 := make([]int,5)
	//切片的声明6
	slice5 := make([]int,5,7)
	//切片的声明7
	slice6 := make([]int,0)

	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(slice4),cap(slice4),slice4)
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(slice5),cap(slice5),slice5)
	fmt.Printf("len=%d,cap=%d,slice=%v",len(slice6),cap(slice6),slice6)
}

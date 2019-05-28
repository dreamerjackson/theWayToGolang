/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//slice 的截取
func sliceTest(){
	numbers:= []int{1,2,3,4,5,6,7,8}
	printSliceInfo(numbers)

	numbers1 :=numbers[1:4]
	printSliceInfo(numbers1)


	numbers2 :=numbers[:3]
	printSliceInfo(numbers2)

	numbers3 :=numbers[3:]
	printSliceInfo(numbers3)
}
//打印切片
func printSliceInfo(x []int){

	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(x),cap(x),x)
}


func main(){

	sliceTest()
}
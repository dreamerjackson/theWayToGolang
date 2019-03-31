/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
//打印
func print(arr []int){

	for  _,data := range arr{
		fmt.Printf("%d ",data)
	}

	fmt.Println()
}


func main(){

	//arr:= []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	arr:=[]int{1,3,2,4,9,2,6,5,4,8}
	print(arr)

	quicksort(arr)

	print(arr)
}

//快速排序
func quicksort(arr []int){

	//print(arr)
	length := len(arr)

	if length <2{
		return
	}
		index:= 0
		start := 0

		for i:=1;i<length;i++{

			if arr[i] <= arr[start]{
				index++
				temp:= arr[index]
				arr[index] = arr[i]
				arr[i] = temp
			}
		}

		tmp:= arr[index]
		arr[index] = arr[start]
		arr[start] = tmp

		quicksort(arr[start:index])
		quicksort(arr[index+1:length])
}
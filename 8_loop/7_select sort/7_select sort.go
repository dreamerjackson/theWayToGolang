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

	arr:= []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	print(arr)

	selectSort(arr)

	print(arr)
}

//选择排序
func selectSort(arr []int){

	length := len(arr)
	//第一个循环从第一个元素到倒数第二个元素。
	for i:= 0;i<length-1;i++{
		//最小的序号
		index:= i
		//遍历其后面的节点，找到最小的节点的下标。
		for j:= i+1;j<length;j++{
			if arr[index]> arr[j]{
				//保留下标
				index = j
			}
		}
		//index != i，就将最小的数交换到a[i]的位置。
		if index != i{
			temp := arr[i]
			arr[i] = arr[index]
			arr[index] = temp
		}
	}
}
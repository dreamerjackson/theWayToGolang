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
	//arr:=[]int{1,3,2,4,9,2,6,5,4,8}
	print(arr)

	insertsort2(arr)

	print(arr)
}

// 插入排序
func insertsort(arr []int){

	//print(arr)
	length := len(arr)

	for i:= 1;i<length;i++{
		temp := arr[i]
		index:=i
		for j:= i-1;j>=0;j--{
			if arr[j] >temp{
				arr[j+1] = arr[j]
			}else{
				break
			}
			index--
		}
		arr[index] = temp
	}
}

//推荐
func insertsort2(arr []int){

	//print(arr)
	length := len(arr)

	for i:= 1;i<length;i++{
		temp := arr[i]
		j:=i-1

		for j>=0 && arr[j] > temp{
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
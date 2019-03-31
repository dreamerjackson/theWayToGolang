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

	//arr:= []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7,6,3,10}
	arr:=[]int{1,3,2,4,9,2,6,5,4,8,0}
	print(arr)
	//println(b)

	shellsort(arr)

	print(arr)
}

//希尔排序
func shellsort(arr []int){

	gap:=4
	length:= len(arr)

	for gap >0{

		for i:=gap;i<length;i+=1{
			j:= i
			temp:= arr[i]
			for;j>0;j-=gap{
				if j-gap>=0 &&  arr[j] < arr[j-gap]{
					tmp:= arr[j]
					arr[j] = arr[j-gap]
					arr[j-gap] = tmp
				}else{
					break
				}
			}
			arr[j] = temp
		}

		gap = gap/2

	}
}
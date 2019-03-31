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

	arr:= []int{2, 6, 9, 3, 1, 4, 0, 7,2, 6, 9, 3, 1,8, 5, 8, 5, 2, 6, 9, 3, 1, 4, 0, 7,2, 6, 9, 3, 1, 4, 0, 7,8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	//arr:=[]int{1,3,2,4,9,2,6,5,4,8}
	print(arr)

	binarysort(arr)

	print(arr)
}

// 插入排序
func binarysort(arr []int){

	//print(arr)
	length := len(arr)
	for i:= 1;i<length;i++{

		tmp:= arr[i]
		 j:= i-1
		 if arr[j] > arr[i]{
		 	 index:= binarysearch(arr,0,i-1,arr[i])

				 for k:= i-1;k>=index;k--{
				 	 arr[k+1] =arr[k]
				 }
				 arr[index] = tmp
		 }
	}
}


func binarysearch(arr []int, low int, high int, data int) int{

	for low <=high{

			mid:= low + (high-low)/2

			if data > arr[mid]{

				if  mid+1< len(arr) && data <=arr[mid+1]{
					return mid+1
				}else{
					low = mid + 1
				}
			}else{
				if mid == 0{
					return 0
				}else{
					high = mid - 1
				}
			}
	}

	return 0

}




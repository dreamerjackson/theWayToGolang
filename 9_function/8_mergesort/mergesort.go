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

	arr = mergeSort(arr)

	print(arr)
}

//归并排序
func mergeSort(arr []int) []int{

	length:= len(arr)

	if length  < 2{
		return arr
	}

	mid:= length/2

	left :=  mergeSort(arr[:mid])
	right:= mergeSort(arr[mid:length])


	return mergeArray(left,right)
}


func mergeArray(left []int, right []int) []int {
	tmp:=make([]int,0)

	i,j:=0,0
	for  i< len(left) && j < len(right){
		if left[i] < right[j]{
			tmp = append(tmp,left[i])
			i++
		}else{
			tmp= append(tmp,right[j])
			j++
		}
	}

	if i==len(left) {
		tmp = append(tmp,right[j:]...)
	}else{
		tmp = append(tmp,left[i:]...)
	}
	return tmp
}
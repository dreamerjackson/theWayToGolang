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

	arr:= []int{1,3,2,4,9,2,6,5,4,8}
	print(arr)

	maopao(arr)

	print(arr)
}

//冒泡排序
func maopao(arr []int){

	length:= len(arr)
	//循环的次数，只用循环 length-1次，即最后只用判断两个数就可以了。
	for i:=0;i<length-1;i++{
		//将数冒到最后，冒完后就不管了，因此每一次需要判断的是length - 1 - i长度。
		for j:=0;j<length-1-i;j++{
			//如果前面的数大于后面的数，交换顺序。
			if arr[j] > arr[j+1]{
				temp:= arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}
//明白了冒泡的原理，就很容易写出一了倒叙的冒泡
func maopao2(arr []int){

	length:= len(arr)
	//和正序一样，倒叙的第一个循环的次数必须是length -1
	for i:=length-1;i>0;i--{
		//第二个循环从最后一个元素开始，一直到只剩下最后一个元素。
		//也就是说，最后一次判断为最后两个数之间的判断。
		for j:= length -1; j>length -1 - i  ;j--{
			//判断
			if arr[j] > arr[j-1]{
				temp:= arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}

		}

	}
}
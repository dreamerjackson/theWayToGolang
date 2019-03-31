/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func main()  {

	arr:= []int{1,3,4,6,7,8,10,13,14}


	for  index,data:= range arr{
		i:= binarySearch(arr,data)
		fmt.Printf("实际序号：%d,找到序号为：%d\n",index,i)

	}

	fmt.Println("-------------------------")
	for  index,data:= range arr{
		i:= binarySearch2(arr,0,len(arr)-1,data)
		fmt.Printf("实际序号：%d,找到序号为：%d\n",index,i)

	}

}



func binarySearch( arr []int, data int ) int{


		low:= 0
		high:= len(arr) - 1

		for low <= high{
			mid:= low + (high-low)/2

			if data > arr[mid]{
				low= mid+1
			}else if data < arr[mid]{
				high= mid - 1
			}else{
				return mid
			}
		}

	return -1
}



func binarySearch2( arr []int,low,high, data int ) int{

		ret:= -1
		if low <= high{
			mid := low + (high-low)/2

			if data > arr[mid]{
				ret = binarySearch2(arr,mid+1,high,data)
			}else if data < arr[mid]{
				ret = binarySearch2(arr,low,mid-1,data)
			}else{
				return mid
			}
		}


	return ret
}
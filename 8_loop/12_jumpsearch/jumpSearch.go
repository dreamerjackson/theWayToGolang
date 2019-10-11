/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	val := jumSearch([]int{1, 2, 10, 25, 34},4, 34)
	fmt.Println(val)
}


func jumSearch(arr []int, sp, value int) int {

	sP := math.Floor(math.Sqrt(float64(sp)))
	pr := 0
	for arr[int(math.Min(sP, float64(len(arr) -1)))] < value  {
		pr := sP
		if arr[len(arr)-1] == value {
			return len(arr)-1
		}
		sP += math.Sqrt(float64(sp))
		if int(pr) >= len(arr) {
			return -1
		}
	}

	for i := pr; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}
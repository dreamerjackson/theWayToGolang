/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"sort"
	"fmt"
)

func main(){
	//对字符串排序
	s:= []string{"a","c","b"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}


func init() {
	//!+ints
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	//!-ints
}
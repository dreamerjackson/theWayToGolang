/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"time"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				a[i]++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}

//
//上面的代码错误的原因是：panic: runtime error: index out of range
//原因是现在的i的结果为10，超过了数组的限制。
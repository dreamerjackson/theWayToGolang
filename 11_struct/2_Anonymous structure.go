/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"math"
	"fmt"
)

func main() {
	//匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)
	fmt.Println(res)

	//匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}
	fmt.Println(addr)

	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}
	fmt.Println(cat)
}

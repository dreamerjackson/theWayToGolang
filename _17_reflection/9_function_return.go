/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

type T struct {
	A, b int
}

func (t T) AddSubThenScale(n int) (int, int) {
	return n * (t.A + t.b), n * (t.A - t.b)
}

func main() {
	t := T{5, 2}
	vt := reflect.ValueOf(t)
	vm := vt.MethodByName("AddSubThenScale")
	results := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(results[0].Int(), results[1].Int()) // 21 9

	neg := func(x int) int {
		return -x
	}
	vf := reflect.ValueOf(neg)
	fmt.Println(vf.Call(results[:1])[0].Int()) // -21
	fmt.Println(vf.Call([]reflect.Value{
		vt.FieldByName("A"), // panic on changing to "b"
	})[0].Int()) // -5
}

//Please note that, non-exported fields shouldn't be used as arguments of reflection calls.
// If the line vt.FieldByName("A") in the above example is replaced with vt.FieldByName("b"), a panic will occur.
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

func main() {
	var s struct {
		X interface{} // an exported field
		y interface{} // a non-exported field
	}
	vp := reflect.ValueOf(&s)
	// If vp represents a pointer. the following
	// line is equivalent to "vs := vp.Elem()".
	vs := reflect.Indirect(vp)
	// vx and vy both represent interface values.
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	// vy is addressable but not modifiable.
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true
	vb := reflect.ValueOf(123)
	vx.Set(vb)     // okay, for vx is modifiable
	// vy.Set(vb)  // will panic, for vy is unmodifiable
	fmt.Println(s) // {123 <nil>}
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true
}
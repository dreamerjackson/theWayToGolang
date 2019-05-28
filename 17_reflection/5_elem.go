/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

/*
we use the method Elem to get the element types of some container types
(a channel type, a map type, a slice type and an array type).
In fact, we can also use this method to get the base type of a pointer
the reflect.Value.Elem method can be also used to get a reflect.Value value which represents the dynamic value of an interface value.
*/
package main

import "fmt"
import "reflect"

func main() {
	var z = 123
	var y = &z
	var x interface{} = y
	v := reflect.ValueOf(&x)
	vx := v.Elem()
	fmt.Println(vx.Kind()) //interface
	vy := vx.Elem()
	fmt.Println(vy.Kind()) // ptr  | this is contain in interface
	vz := vy.Elem()
	fmt.Println(vz.Kind()) // int  |  this is the ptr's basic type
	vz.Set(reflect.ValueOf(789))
	fmt.Println(z) // 789
}
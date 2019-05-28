/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"reflect"
)



/*
The concrete type of interface{} is represented by
reflect.Type and the underlying value is represented by reflect.Value.
There are two functions reflect.TypeOf() and reflect.ValueOf()
which return the reflect.Type and reflect.Value respectively.
These two types are the base to create our query generator.
Let's write a simple example to understand these two types.

*/
func main() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))//type:  float64
	fmt.Println("value: ", reflect.ValueOf(num))//value:  1.2345
	fmt.Println("value:", reflect.ValueOf(num).String())//value: <float64 Value>

	fmt.Println("type:", reflect.ValueOf(num).Type())//type: float64
	fmt.Println("kind is float64: ", reflect.TypeOf(num).Kind() == reflect.Float64)//kind is float64: true
	fmt.Println("kind is float64:", reflect.ValueOf(num).Kind() == reflect.Float64)//kind is float64: true
	fmt.Println("value:", reflect.ValueOf(num).Float())//value: 1.2345


	type order struct {
		ordId      int
		customerId int
	}
	o := order{
		ordId:      456,
		customerId: 56,
	}
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	fmt.Println("Type ", t)//Type  main.order
	fmt.Println("Value ", v)//Value  {456 56}
}









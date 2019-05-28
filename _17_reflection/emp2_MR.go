package main

import (
	"fmt"
	"reflect"
)

func Map(slice interface{}, fn func(a interface{}) interface{}) interface{} {
	val := reflect.ValueOf(slice)
	out := reflect.MakeSlice(reflect.TypeOf(slice), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		ret := fn(val.Index(i).Interface())
		out.Index(i).Set(reflect.ValueOf(ret))
	}
	return out.Interface()
}

func main() {
	a := []int{1, 2, 3}
	fn := func(a interface{}) interface{} {
		return a.(int) * 2
	}
	b := Map(a, fn)
	fmt.Printf("%v\n", b)
}

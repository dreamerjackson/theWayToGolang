
package main

import "reflect"
//接收任意类型slices,并统一返回[]interface{}
func handleSlice(arg interface{}) (out []interface{}, ok bool) {
	argValue := reflect.ValueOf(arg)
	if argValue.Type().Kind() == reflect.Slice {
		length := argValue.Len()
		if length == 0 {
			return
		}
		ok = true
		out = make([]interface{}, length)
		for i := 0; i < length; i++ {
			out[i] = argValue.Index(i).Interface()
		}
	}
	return
}

/*
注意下面的方法是错误的： 因为例如[]int 是不能赋值给[]interface的
func handleSlice(in []interface{}) (out []interface{}) {
	out = make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return
}
*/
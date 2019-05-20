package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"gopl.io/ch12/format"
)


//基于反射将可以处理任何的类型
// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	//!+time
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))                  // "1"
	fmt.Println(format.Any(d))                  // "1"
	fmt.Println(format.Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"

}

/*
到目前为止, 我们的函数将每个值视作一个不可分割没有内部结构的物品, 因此它叫 formatAtom.
对于聚合类型(结构体和数组)和接口，只是打印值的类型,
对于引用类型(channels, functions, pointers, slices, 和 maps), 打印类型和十六进制的引用地址.
虽然还不够理想, 但是依然是一个重大的进步, 并且 Kind 只关心底层表示, format.Any 也支持具名类型. 例如:

*/
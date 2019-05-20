package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
	fmt.Println("value:", reflect.ValueOf(num).String())

	fmt.Println("type:", reflect.ValueOf(num).Type())
	fmt.Println("kind is float64: ", reflect.TypeOf(num).Kind() == reflect.Float64)
	fmt.Println("kind is float64:", reflect.ValueOf(num).Kind() == reflect.Float64)
	fmt.Println("value:", reflect.ValueOf(num).Float())

}

//运行结果:
//type:  float64
//value:  1.2345
//value: <float64 Value>
//type: float64
//kind is float64: true
//kind is float64: true
//value: 1.2345
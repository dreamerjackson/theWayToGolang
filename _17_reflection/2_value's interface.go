

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

//ValueOf can change to  interface{}
package main

import (
"fmt"
"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	fmt.Println(pointer.Interface())
	fmt.Println(value.Interface())

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

//运行结果：
//0xc000012080
//1.2345
//0xc000012080
//1.2345

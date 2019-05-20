

//反射转换到接口


//inerface可以直接放入fmt中打印
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

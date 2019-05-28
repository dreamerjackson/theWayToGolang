/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

func main() {
	c := make(chan string, 2)
	vc := reflect.ValueOf(c)
	vc.Send(reflect.ValueOf("C"))
	succeeded := vc.TrySend(reflect.ValueOf("Go"))
	fmt.Println(succeeded) // true
	succeeded = vc.TrySend(reflect.ValueOf("C++"))
	fmt.Println(succeeded) // false
	fmt.Println(vc.Len(), vc.Cap()) // 2 2
	vs, succeeded := vc.TryRecv()
	fmt.Println(vs.String(), succeeded) // C true
	vs, sentBeforeClosed := vc.Recv()
	fmt.Println(vs.String(), sentBeforeClosed) // Go false
	vs, succeeded = vc.TryRecv()
	fmt.Println(vs.String()) // <invalid Value>
	fmt.Println(succeeded)   // false
}
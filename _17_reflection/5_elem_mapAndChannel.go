/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

func main() {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	// The next line prints: map array slice
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())
	tx, ty := ta.Elem(), tb.Elem()

	// byte is an alias of uint8
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8
	fmt.Println(tx.Bits(), ty.Bits()) // 16 8
	fmt.Println(tx.ConvertibleTo(ty)) // true
	fmt.Println(tb.ConvertibleTo(ta)) // false

	// Slice and map types are uncomparable.
	fmt.Println(tb.Comparable()) // false
	fmt.Println(tm.Comparable()) // false
	fmt.Println(ta.Comparable()) // true
	fmt.Println(tc.Comparable()) // true
}
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

func main() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeeded := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, vc.Len(), vc.Cap()) // true 1 1

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vZero, Send: vZero},
		{Dir: reflect.SelectRecv, Chan: vc, Send: vZero},
		{Dir: reflect.SelectSend, Chan: vc, Send: vSend},
	}
	selIndex, vRecv, closed := reflect.Select(branches)
	fmt.Println(selIndex)    // 1
	fmt.Println(closed)      // true
	fmt.Println(vRecv.Int()) // 123
	vc.Close()
	// Remove the send case branch this time,
	// for it may cause panic.
	selIndex, _, closed = reflect.Select(branches[:2])
	fmt.Println(selIndex, closed) // 1 false
}
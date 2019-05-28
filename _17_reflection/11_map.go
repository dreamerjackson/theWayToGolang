/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
import "reflect"

//
func main() {
	valueOf := reflect.ValueOf
	m := map[string]int{"Unix": 1973, "Windows": 1985}
	v := valueOf(m)
	// A zero second Value argument means to delete an entry.
	v.SetMapIndex(valueOf("Windows"), reflect.Value{})
	v.SetMapIndex(valueOf("Linux"), valueOf(1991))
	//Please note that, the MapRange method is supported since Go 1.12.
	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t:", i.Value())
	}
}
//result:
// Unix 	: 1973
// Linux 	: 1991
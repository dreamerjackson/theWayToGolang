/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t) //main.order
	fmt.Println("Kind ", k) //struct
}

/*
There is one more important type in the reflection package called Kind.

The types Kind and Type in the reflection package might seem similar
but they have a difference which will be clear from the program below.
*/
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}

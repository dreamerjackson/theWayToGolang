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

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}
/*
运行结果：
old value of pointer: 1.2345
type of pointer: float64
settability of pointer: true
new value of pointer: 77



*/

/*
reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。
x := 12
a := reflect.ValueOf(2)
println(a.CanAddr()) //false

b := reflect.ValueOf(x)
println(b.CanAddr()) //false

c := reflect.ValueOf(&x)
println(c.CanAddr()) //false

d := reflect.ValueOf(&x).Elem()
println(d.CanAddr()) //true

a不可取址，因为a中的值仅仅是12的拷贝的副本

b不可取址，因为b中的值是a的拷贝的副本

c不可取址，因为c中的值只是一个指针&x的拷贝

d可取址，因为Elem()方法，可以获取变量对应的可取址地址的value
*/

/*
example2:

package main

import "fmt"
import "reflect"

func main() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem() // get the value referenced by vp
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789)) // <=> vn.SetInt(789)
	fmt.Println(n)               // 789
}


*/

/*
example3:
package main

import "fmt"
import "reflect"

func main() {
	var s struct {
		X interface{} // an exported field
		y interface{} // a non-exported field
	}
	vp := reflect.ValueOf(&s)
	// If vp represents a pointer. the following
	// line is equivalent to "vs := vp.Elem()".
	vs := reflect.Indirect(vp)
	// vx and vy both represent interface values.
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	// vy is addressable but not modifiable.
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true
	vb := reflect.ValueOf(123)
	vx.Set(vb)     // okay, for vx is modifiable
	// vy.Set(vb)  // will panic, for vy is unmodifiable
	fmt.Println(s) // {123 <nil>}
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true
}
*/
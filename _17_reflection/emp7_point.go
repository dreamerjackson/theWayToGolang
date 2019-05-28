/*
 * Copyright (c) 2019. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

var handler = func(u *User, message string) {
	fmt.Printf("Hello, My name is %s, I am %d years old ! so, %s\n", u.Name, u.Age, message)
}

//使用普通反射的方式处理名字屏蔽
func filtName(u *User, message string) {
	fn := reflect.ValueOf(handler)
	uv := reflect.ValueOf(u)
	name := uv.Elem().FieldByName("Name")
	name.SetString("XXX")
	fn.Call([]reflect.Value{uv, reflect.ValueOf(message)})
}


//重用部分数据减少重复创建的反射方式处理名字屏蔽
var offset uintptr = 0xFFFF
func filtNameWithReuseOffset(u *User, message string) {
	if offset == 0xFFFF {
		t := reflect.TypeOf(u).Elem()
		name, _ := t.FieldByName("Name")
		offset = name.Offset

	}
	up := (*[2]uintptr)(unsafe.Pointer(&u))
	upnamePtr := (*string)(unsafe.Pointer(up[0] + offset))
	*upnamePtr = "YYY"
	fn := reflect.ValueOf(handler)
	uv := reflect.ValueOf(u)
	fn.Call([]reflect.Value{uv, reflect.ValueOf(message)})
}

func main(){
	filtNameWithReuseOffset(&User{Name:"jsonson",Age:69},"hello world")
}
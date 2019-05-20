
package main

import (
	"sync"
	"reflect"
	"fmt"
)



type JunkServer struct {
	mu   sync.Mutex
	log1 []string
	log2 []int
}

func  Handler1(args int, reply *int) {
	*reply = args
}


func  Handler2(args int, reply *int) {
	*reply = args
}

func main(){
	v1:= reflect.ValueOf(Handler1)


	//int类型
	str:= 123
	s:= reflect.ValueOf(str)

	//int指针类型
	replyv := reflect.New(reflect.TypeOf(2))

	v1.Call([]reflect.Value{s,replyv})
	fmt.Println(replyv.Elem())


	//指针类型
	args := reflect.New(reflect.TypeOf(2))
	//a:= args.Interface().(*int)
	//*a = 456

	v2:= reflect.ValueOf(Handler2)
	//设置指针类型的值
	v2.Call([]reflect.Value{args.Elem(),replyv})
	fmt.Println(replyv.Elem())
}


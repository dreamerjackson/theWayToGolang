/*
 * Copyright (c) 2019. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
)


type stu struct{
	Name string
	Grade int
}

func (s stu)print(){

}

type face interface {
	print()
}
func main(){

	var f = func(c context.Context,s *stu) {
		fmt.Printf("%+v",*s)
	}

	WithProto(f)

}
func WithProto(f interface{}) {
	typeOfProto := reflect.TypeOf((*face)(nil)).Elem()
	typeOfContext := reflect.TypeOf((*context.Context)(nil)).Elem()

	// func sign
	if f == nil {
		log.Fatalf("WithCallback : f is nil")
	}
	typeOfFunc := reflect.TypeOf(f)
	if typeOfFunc.Kind() != reflect.Func {
		log.Fatalf("WithCallback : f is not a function")
	}
	if typeOfFunc.NumIn() != 2 {
		log.Fatalf("WithCallback : f parameter size is not 2")
	}
	if typeOfFunc.NumOut() != 0 {
		log.Fatalf("WithCallback : f return size is not 0")
	}

	// context
	firstTypeOfFunc := typeOfFunc.In(0)
	if firstTypeOfFunc != typeOfContext {
		log.Fatalf("WithCallback : f 1st parameter is not context.Context")
	}

	// proto.message
	secondTypeOfFunc := typeOfFunc.In(1)
	if secondTypeOfFunc.Kind() != reflect.Ptr {
		log.Fatalf("WithCallback : f 2nd parameter is not a pointer")
	}

	if secondTypeOfFunc.Elem().Kind() != reflect.Struct {
		log.Fatalf("WithCallback : f 2nd parameter is not a struct pointer")
	}

	if !secondTypeOfFunc.Implements(typeOfProto) {
		log.Fatalf("WithCallback : f 2nd parameter is not implements proto.Message")
	}

	ctx := context.Background()
	body := reflect.New(secondTypeOfFunc.Elem())


	var s stu = stu{
		Name:"jonson",
		Grade:100,
	}
	k:= reflect.ValueOf(s)
	body.Elem().Set(k) //指针的内容赋值
	reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(ctx), body})
}


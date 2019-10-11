package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

type Middleware func(f http.HandlerFunc) http.HandlerFunc

func advanceworking(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("正在处理...",time.Now())
		f(w,r)
	}
}



func advancemullogging(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w,r)
	}
}

func advancemulhello( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello jonson")
}

func advancemulworld( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello olaya")
}

func Chain(f http.HandlerFunc,middleware ...Middleware)  http.HandlerFunc{

	for _,m := range middleware{
		f = m(f)
	}

	return f
}


func main(){

	server:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/hello",Chain(advancemulhello,advancemullogging,advanceworking))
	http.HandleFunc("/world",Chain(advancemulworld,advancemullogging,advanceworking))

	server.ListenAndServe()
}
package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)


func working(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("正在处理...",time.Now())
		f(w,r)
	}
}



func mullogging(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w,r)
	}
}

func mulhello( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello jonson")
}

func mulworld( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello olaya")
}

func main(){

	server:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/hello",working(mullogging(mulhello)))
	http.HandleFunc("/world",working(mullogging(mulworld)))

	server.ListenAndServe()
}
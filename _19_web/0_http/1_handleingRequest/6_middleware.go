package main

import (
	"net/http"
	"fmt"
	"log"
)


func logging(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w,r)
	}
}

func middlehello( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello jonson")
}

func middleworld( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello olaya")
}

func main(){

	server:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/hello",logging(middlehello))
	http.HandleFunc("/world",logging(middleworld))

	server.ListenAndServe()
}
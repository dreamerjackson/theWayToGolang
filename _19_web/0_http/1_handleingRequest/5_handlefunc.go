package main

import (
	"net/http"
	"fmt"
)


func hello( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello jonson")
}

func world( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello olaya")
}

func main(){

	server:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/hello",hello)
	http.HandleFunc("/world",world)

	server.ListenAndServe()
}
package main

import (
	"net/http"
	"fmt"
)



type MyHandler struct {}

func (m MyHandler) ServeHTTP( w http.ResponseWriter,  r *http.Request) {
	 fmt.Fprintf(w,"hello world")
}

func main(){
	handler :=MyHandler{}
	server:= http.Server{
		Addr:"127.0.0.1:8080",
		Handler:handler,
	}

	server.ListenAndServe()
}
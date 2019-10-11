package main

import (
	"net/http"
	"fmt"
)



type Hello struct {}

func (m Hello) ServeHTTP( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello jonson")
}

type World struct {}

func (m World) ServeHTTP( w http.ResponseWriter,  r *http.Request) {
	fmt.Fprintf(w,"hello olaya")
}

func main(){
	hello :=Hello{}
	world:= World{}
	server:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.Handle("/hello",hello)
	http.Handle("/world",world)

	server.ListenAndServe()
}
package main

import (
	"net/http"
	"fmt"
)


func process3(w http.ResponseWriter,r * http.Request){

	length := r.ContentLength

	body:= make([]byte,length)

	r.Body.Read(body)

	fmt.Fprintln(w,string(body))
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process3)

	serve.ListenAndServe()
}

package main

import (
	"net/http"
	"fmt"
)


func writeHeader(w http.ResponseWriter,r * http.Request){

	w.WriteHeader(501)

	fmt.Fprintln(w,"no such servier")


}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/writeHeader",writeHeader)

	serve.ListenAndServe()
}

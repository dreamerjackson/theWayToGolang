package main

import (
	"net/http"
	"fmt"
)


func process2(w http.ResponseWriter,r * http.Request){

	v := r.URL.Query();
	a := v.Get("a")
	fmt.Println("a:",a)
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process2)

	serve.ListenAndServe()
}

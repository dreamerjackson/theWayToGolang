package main

import (
	"net/http"
	"fmt"
)


func process4(w http.ResponseWriter,r * http.Request){

	r.ParseForm()
	fmt.Fprintln(w,r.Form.Get("name"))
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process4)

	serve.ListenAndServe()
}

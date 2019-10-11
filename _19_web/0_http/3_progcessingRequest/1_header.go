package main

import (
	"net/http"
	"fmt"
)


func process(w http.ResponseWriter,r * http.Request){

	 head:= r.Header

	fmt.Fprintln(w,head)

}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process)

	serve.ListenAndServe()
}

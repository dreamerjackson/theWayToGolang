package main

import (
	"net/http"
)


func writeRedirect(w http.ResponseWriter,r * http.Request){


	w.Header().Set("Location","http://baidu.com")
	w.WriteHeader(302)
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/writeRedirect",writeRedirect)

	serve.ListenAndServe()
}

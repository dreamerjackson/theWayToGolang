package main

import (
	"net/http"
	"fmt"
)

//只会解析post当中的数据
func process5(w http.ResponseWriter,r * http.Request){

	r.ParseForm()
	fmt.Fprintln(w,r.PostForm)
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process5)

	serve.ListenAndServe()
}

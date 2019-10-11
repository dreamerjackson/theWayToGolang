package main

import (
	"net/http"
	"encoding/json"
)



type Post struct{
	User string
	Age int
}

func writejonson(w http.ResponseWriter,r * http.Request){


	w.Header().Set("Conten-Type","application/json")
	post:= Post{
		User:"josnon",
		Age:30,
	}

	json,_:= json.Marshal(post)

	w.Write(json)

}


func main(){
	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/writejonson",writejonson)
	serve.ListenAndServe()
}

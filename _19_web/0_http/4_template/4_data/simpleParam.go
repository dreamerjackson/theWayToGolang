package main

import (
	"net/http"
	"text/template"
	"log"
)


func tpl(w http.ResponseWriter,r * http.Request){
	templ,err:= template.ParseFiles("tpl.html")
	if err!=nil{
		log.Fatalln(err)
	}

	templ.Execute(w,"olaya")
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/tmp",tpl)

	serve.ListenAndServe()
}

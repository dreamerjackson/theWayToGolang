package main

import (
	"net/http"
)


func write(w http.ResponseWriter,r * http.Request){

	str:=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>go web programing</title>
</head>
<body>
<form method="post" action="http://localhost:8080/process?hello=world&name=jonson">
    <div>Hello world</div>

</form>
</body>
</html>`
	//fmt.Fprintln(w,str)
	w.Write([]byte(str))
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/write",write)

	serve.ListenAndServe()
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)


func process6(w http.ResponseWriter,r * http.Request){
	//fmt.Fprintln(w,r.FormValue("name"))
	fmt.Println(123)

	s:=r.PostFormValue("post")
	body, _ := ioutil.ReadAll(r.Body)
	type TopUserTaskRequestModel struct {
		TimeRange int
		Type      int //  0 loss  1 profit
	}

	var requestModel TopUserTaskRequestModel
	if err := json.Unmarshal(body, &requestModel); err != nil {
fmt.Println(err)	}

	fmt.Println("requestmodel:",requestModel)
	fmt.Fprintln(w,s)//只会解析post当中的数据
}


func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/process",process6)

	serve.ListenAndServe()
}

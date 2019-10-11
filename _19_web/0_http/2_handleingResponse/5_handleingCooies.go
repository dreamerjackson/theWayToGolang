package main

import (
	"net/http"
	"time"
	"fmt"
)



//更简便的设置cookie的方式
func setCookiequick(w http.ResponseWriter,r * http.Request){

	c1:=http.Cookie{
		Name:"third_cookie",
		Value:"jonson",
	}

	c2:=http.Cookie{
		Name:"four_cookie",
		Value:"olaya",
		Expires:time.Now().Add(1*time.Hour),
	}

	http.SetCookie(w,&c1)
	http.SetCookie(w,&c2)


}

func setCookie(w http.ResponseWriter,r * http.Request){

	c1:=http.Cookie{
		Name:"first_cookie",
		Value:"jonson",
	}

	c2:=http.Cookie{
		Name:"second_cookie",
		Value:"olaya",
		Expires:time.Now().Add(1*time.Hour),
	}

	w.Header().Set("Set-Cookie",c1.String())
	w.Header().Add("Set-Cookie",c2.String())

}



func getcookie(w http.ResponseWriter,r * http.Request){
		//h:=r.Header["Cookie"]
		//fmt.Fprintln(w,h)


	c1,err:=	r.Cookie("first_cookie")

	if err!=nil{
		fmt.Fprintln(w,"cannot get the first-cookie")
	}

	cs:= r.Cookies()

	fmt.Fprintln(w,c1)
	fmt.Fprintln(w,cs)



}



func main(){


	serve:= http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/setCookie",setCookie)
	http.HandleFunc("/setCookiequick",setCookiequick)
	http.HandleFunc("/getcookie",getcookie)

	serve.ListenAndServe()
}

package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
)


//解析多个文件
func main() {

	t, err := template.ParseFiles("one.demo", "two.demo", "three.demo")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
	err = t.ExecuteTemplate(os.Stdout,"two.demo",nil)

	if err != nil {
		log.Fatalln(err)
	}
}

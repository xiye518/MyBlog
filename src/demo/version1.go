package main

import (
	"net/http"
	"io"
	"log"
)

//Version 1
func main(){
	//设置路由
	http.HandleFunc("/",sayHello)

	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}

}

func sayHello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello World,this is virson 1.")
}

//Version 2
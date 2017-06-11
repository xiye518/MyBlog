package main

import (
	"net/http"
	"io"
	"log"
)

//通过mux控制访问
//可以修改为静态文件服务器
func main(){
	mux :=http.NewServeMux()

	mux.Handle("/",&virsion2Handler{})
	mux.HandleFunc("/hello",sayHello2)

	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		log.Fatal(err)
	}
}

type virsion2Handler struct{}


func (*virsion2Handler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"url:"+r.URL.String())
}

func sayHello2(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello World,this is virsion 2.")
}
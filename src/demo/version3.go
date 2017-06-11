package main

import (
	"net/http"
	"time"
	"log"
	"io"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main(){
	server := http.Server{
		Addr:":8080",
		Handler:&virsion3Handler{},
		ReadTimeout:5*time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello3
	mux["/bye"] = sayBye

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

type virsion3Handler struct{}

func (*virsion3Handler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if h,ok :=mux[r.URL.String()];ok{
		h(w,r)
		return
	}

	io.WriteString(w,"url:"+r.URL.String())
}

func sayHello3(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello World,this is virsion 3.")
}

func sayBye(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Bye bye,this is virson 3.")
}

package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hello World</h2>")
}

func check(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Health Check</h1>")
}

func main(){
	http.HandleFunc("/",index)
	http.HandleFunc("/health",check)
	fmt.Printf("server running")
	http.ListenAndServe(":3000",nil)
}


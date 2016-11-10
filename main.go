package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	var x = r.Referer()
	fmt.Fprintf(w, "Hello Mark\n")
	fmt.Fprintf(w, x)
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Earth\n")
}
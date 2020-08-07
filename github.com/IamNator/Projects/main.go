package main

import ("fmt"
		"net/http")


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
}

func main(){
	http.HandleFunc("/", handler)
	defer http.ListenAndServe(":8080", nil)
	fmt.Println("started")
}
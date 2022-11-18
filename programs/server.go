package main

import (
	"fmt"
	"net/http"
)



func helloHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, aj!")
}

func main() {
	
	http.HandleFunc("/hello", helloHandleFunc)

	http.ListenAndServe("localhost:8080", nil) 

}



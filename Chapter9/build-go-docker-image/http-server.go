package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! ")
}

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error starting http server :: ", err)
	}
}

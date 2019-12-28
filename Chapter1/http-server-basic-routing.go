package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World !")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login Page")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout Page")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("errro starting http server: ", err)
		return
	}
}

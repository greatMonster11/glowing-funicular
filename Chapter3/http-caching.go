package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var newCache *cache.Cache

func init() {
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	newCache.Set("foo", "bar", cache.DefaultExpiration)
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	foo, found := newCache.Get("foo")
	if found {
		log.Printf("Key found in cache with value as :: ", foo.(string))
		fmt.Fprintf(w, "Hello "+foo.(string))
	} else {
		log.Printf("Key not found in cache :: ", "foo")
		fmt.Fprintf(w, "Key not found in cache")
	}
}

func main() {
	http.HandleFunc("/", getFromCache)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

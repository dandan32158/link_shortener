package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	mu.Lock()
	urlStore[shortURL] = longURL
	mu.Unlock()

	fmt.Fprintf(w, "Shortened URL: /%s\n", shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	mu.Lock()
	longURL, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

package main

import (
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(getEnv("HELLO_PORT", ":8080"), nil); err != nil {
		panic(err)
	}
}

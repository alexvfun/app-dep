package main

import (
	"fmt"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	fmt.Println("FOO:", os.Getenv("HELLO_PORT"))
	if err := http.ListenAndServe(os.Getenv("HELLO_PORT"), nil); err != nil {
		panic(err)
	}
}

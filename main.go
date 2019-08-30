package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "hello: ", log.Llongfile|log.Lmicroseconds|log.LstdFlags)
	logger.Println("starting...")

	addr, set := os.LookupEnv("ADDR")
	if !set {
		addr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("hit")
		fmt.Fprintln(w, "hello")
	})

	logger.Println("serving on", addr)
	logger.Fatal(http.ListenAndServe(addr, mux))
}

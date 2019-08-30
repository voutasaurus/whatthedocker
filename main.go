package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr, set := os.LookupEnv("ADDR")
	if !set {
		addr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	log.Fatal(http.ListenAndServe(addr, mux))
}

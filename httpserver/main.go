package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	 fmt.Printf("/ request\n")
	io.WriteString(w, "Website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("/hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

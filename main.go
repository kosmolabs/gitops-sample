package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	t := time.Now()
	fmt.Fprintf(w, "Hello there! time is: %s\n", t)
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)
}

func main() {
	// Register the handler function for the root ("/") route
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

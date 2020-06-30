package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Full Cycle by FZANATA")
}

func main() {
	fmt.Println("Hello Full Cycle")

	port := 8080
	fmt.Printf("Run with port mapping: -p %d:%d\n", port, port)

	http.HandleFunc("/", greet)

	fmt.Printf("Server running on http://localhost:%d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

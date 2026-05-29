package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"globalfin-backend"}`)
	})

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
// Test final de pipeline1

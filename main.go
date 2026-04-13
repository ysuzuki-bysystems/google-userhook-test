package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s\n", r.URL)
		fmt.Printf("%#v\n", r.Header)

		w.WriteHeader(200)
	})
	http.ListenAndServe(addr, nil)
}

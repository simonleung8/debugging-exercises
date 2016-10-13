package main

import (
	"fmt"
	"net/http"
	"os"
)

func bigBytes() *[]byte {
	s := make([]byte, 50000000)
	return &s
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, app is running")
}

func main() {
	port := os.Getenv("PORT")
	fmt.Fprintf(os.Stdout, "Obtained port from Env $PORT: "+port)

	fmt.Fprintf(os.Stderr, "App requires minimum 256M to run, make sure enough memory is allocated")
	for i := 0; i < 3; i++ {
		s := bigBytes()
		if s == nil {
			fmt.Fprintf(os.Stderr, "App run out of memory, 256M is required")
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

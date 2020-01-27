package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Process request, URL %q, Method %s\n", r.URL, r.Method)

		_, _ = fmt.Fprint(w, "Hello World")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

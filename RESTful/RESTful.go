package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	content := "Hello ALaa TO GO LAND Progreming , %q"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, content, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

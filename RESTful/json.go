package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	content := "Hello ALaa TO GO LAND Progreming AAAAAAAA"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(content)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

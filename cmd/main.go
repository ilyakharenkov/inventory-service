package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/products/stock", func(w http.ResponseWriter, r *http.Request) {})
}

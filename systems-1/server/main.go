package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  http.HandleFunc("/api", handler)
  http.ListenAndServe(":3333", nil)
}
